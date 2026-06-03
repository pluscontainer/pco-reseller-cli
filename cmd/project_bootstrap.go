/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
)

var bootstrapUserName, bootstrapPassword, bootstrapDescription string
var bootstrapWithDefaultNetwork, bootstrapWithDefaultQuota bool

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

		pw := bootstrapPassword
		if !cmd.Flags().Changed("password") {
			var err error
			pw, err = password.Generate(24, 4, 4, false, true)
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
		fmt.Printf("Created project %s (%s)\n", project.Name, project.Id)

		user, err := psOsClient.CreateUser(ctx, openapi.CreateOpenStackUser{
			Name:           userName,
			Description:    description,
			Enabled:        &enabled,
			DefaultProject: &project.Id,
			Password:       pw,
		})
		if err != nil {
			return fmt.Errorf("error creating user: %w", err)
		}
		fmt.Printf("Created user %s (%s)\n", user.Name, user.Id)

		if err := psOsClient.AddUserToProject(ctx, project.Id, user.Id); err != nil {
			return fmt.Errorf("error adding user to project: %w", err)
		}
		fmt.Printf("Assigned user %s to project %s\n", user.Name, project.Name)

		if bootstrapWithDefaultQuota {
			if _, err := psOsClient.UpdateProjectQuota(ctx, project.Id, defaultQuota); err != nil {
				return fmt.Errorf("error applying default quota: %w", err)
			}
			fmt.Printf("Applied default quota to project %s\n", project.Name)
		}

		fmt.Println()
		fmt.Println("Bootstrap completed successfully")
		fmt.Println("--------------------------------")
		fmt.Printf("Project Name: %s\n", project.Name)
		fmt.Printf("Project ID:   %s\n", project.Id)
		fmt.Printf("User Name:    %s\n", user.Name)
		fmt.Printf("Password:     %s\n", pw)
		return nil
	},
}

func init() {
	projectCmd.AddCommand(bootstrapCmd)
	bootstrapCmd.Flags().StringVarP(&bootstrapDescription, "description", "d", "", "Description of the project and user (defaults to the project name if not set)")
	bootstrapCmd.Flags().StringVar(&bootstrapUserName, "user-name", "", "Name of the user to create (defaults to <project-name>-admin)")
	bootstrapCmd.Flags().StringVarP(&bootstrapPassword, "password", "p", "", "Password for the new user (auto-generated if not set)")
	bootstrapCmd.Flags().BoolVar(&bootstrapWithDefaultNetwork, "with-default-network", false, "Preconfigure the project with a default network, router and security groups")
	bootstrapCmd.Flags().BoolVar(&bootstrapWithDefaultQuota, "with-default-quota", false, "Apply the standard quota set to the project")
}
