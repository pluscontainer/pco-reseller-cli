/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var userDescription, userDefaultProject, userPassword string

var userCreateCmd = &cobra.Command{
	Use:   "create [user-name]",
	Short: "Create a new reseller user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		enabled := true

		description := userDescription
		if !cmd.Flags().Changed("description") {
			description = args[0]
		}

		resp, err := psOsClient.CreateUser(ctx, openapi.CreateOpenStackUser{
			Name:           args[0],
			Description:    description,
			Enabled:        &enabled,
			DefaultProject: &userDefaultProject,
			Password:       userPassword,
		})
		if err != nil {
			return err
		}

		printUsers([]openapi.CreatedOpenStackUser{*resp})
		return nil
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)
	userCreateCmd.Flags().StringVarP(&userDescription, "description", "d", "", "Description of the user (defaults to the user name if not set)")
	userCreateCmd.Flags().StringVar(&userDefaultProject, "default-project", "", "Default project of the user")
	if err := userCreateCmd.MarkFlagRequired("default-project"); err != nil {
		panic(err)
	}
	userCreateCmd.Flags().StringVarP(&userPassword, "password", "p", "", "Password of the user")
	if err := userCreateCmd.MarkFlagRequired("password"); err != nil {
		panic(err)
	}
}
