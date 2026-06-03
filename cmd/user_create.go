/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var userDescription, userDefaultProject, userPassword string

// createCmd represents the create command
var userCreateCmd = &cobra.Command{
	Use:   "create [user-name]",
	Short: "Create a new reseller user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		enabled := true

		resp, err := psOsClient.CreateUser(ctx, openapi.CreateOpenStackUser{
			Name:           args[0],
			Description:    userDescription,
			Enabled:        &enabled,
			DefaultProject: &userDefaultProject,
			Password:       userPassword,
		})

		if err != nil {
			return err
		}

		fmt.Println(resp.Id)
		return nil
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)

	userCreateCmd.Flags().StringVarP(&userDescription, "description", "d", "No Description", "Specify the description of the user")

	userCreateCmd.Flags().StringVar(&userDefaultProject, "default-project", "", "Specify the default project of the user")
	if err := userCreateCmd.MarkFlagRequired("default-project"); err != nil {
		panic(err)
	}

	userCreateCmd.Flags().StringVarP(&userPassword, "password", "p", "", "Specify the password of the user")
	if err := userCreateCmd.MarkFlagRequired("password"); err != nil {
		panic(err)
	}
}
