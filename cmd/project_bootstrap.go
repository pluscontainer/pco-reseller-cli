/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var bootstrapUserName, bootstrapPassword, bootstrapDescription string
var bootstrapWithDefaultNetwork bool

const passwordChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#%^&*-_"

func generatePassword(length int) (string, error) {
	buf := make([]byte, length)
	for i := range buf {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(passwordChars))))
		if err != nil {
			return "", err
		}
		buf[i] = passwordChars[n.Int64()]
	}
	return string(buf), nil
}

var bootstrapCmd = &cobra.Command{
	Use:     "bootstrap [project-name]",
	Short:   "Bootstrap a project with a user",
	Long:    `Creates a project, a user and assigns the user to the project in one step.`,
	Example: "  pco-reseller-cli project bootstrap my-project\n  pco-reseller-cli project bootstrap my-project --password secret\n  pco-reseller-cli project bootstrap my-project --user-name my-user --with-default-network",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		description := bootstrapDescription
		if !cmd.Flags().Changed("description") {
			description = projectName
		}

		userName := bootstrapUserName
		if !cmd.Flags().Changed("user-name") {
			userName = projectName + "-admin"
		}

		password := bootstrapPassword
		if !cmd.Flags().Changed("password") {
			var err error
			password, err = generatePassword(24)
			if err != nil {
				return fmt.Errorf("error generating password: %w", err)
			}
		}

		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		enabled := true

		project, err := psOsClient.CreateProject(ctx, openapi.ProjectCreate{
			Name:                projectName,
			Description:         description,
			Enabled:             &enabled,
			NetworkPreconfigure: &bootstrapWithDefaultNetwork,
		})
		if err != nil {
			return fmt.Errorf("error creating project: %w", err)
		}

		user, err := psOsClient.CreateUser(ctx, openapi.CreateOpenStackUser{
			Name:           userName,
			Description:    description,
			Enabled:        &enabled,
			DefaultProject: &project.Id,
			Password:       password,
		})
		if err != nil {
			return fmt.Errorf("error creating user: %w", err)
		}

		if err := psOsClient.AddUserToProject(ctx, project.Id, user.Id); err != nil {
			return fmt.Errorf("error adding user to project: %w", err)
		}

		fmt.Println()
		fmt.Println("Bootstrap completed successfully")
		fmt.Println("--------------------------------")
		fmt.Printf("Project Name: %s\n", project.Name)
		fmt.Printf("Project ID:   %s\n", project.Id)
		fmt.Printf("User Name:    %s\n", user.Name)
		fmt.Printf("Password:     %s\n", password)
		return nil
	},
}

func init() {
	projectCmd.AddCommand(bootstrapCmd)

	bootstrapCmd.Flags().StringVarP(&bootstrapDescription, "description", "d", "", "Description of the project and user (defaults to the project name if not set)")
	bootstrapCmd.Flags().StringVar(&bootstrapUserName, "user-name", "", "Name of the user to create (defaults to <project-name>-admin)")
	bootstrapCmd.Flags().StringVarP(&bootstrapPassword, "password", "p", "", "Password for the new user (auto-generated if not set)")
	bootstrapCmd.Flags().BoolVar(&bootstrapWithDefaultNetwork, "with-default-network", false, "Preconfigure the project with a default network, router and security groups")
}
