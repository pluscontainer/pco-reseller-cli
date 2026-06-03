/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var updateUserName, updateUserDescription, updateUserDefaultProject, updateUserPassword string

var enableUser, disableUser bool

// createCmd represents the create command
var userUpdateCmd = &cobra.Command{
	Use:   "update [user-id]",
	Short: "Update a reseller user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()

		resp, err := psOsClient.GetUser(ctx, args[0])
		if err != nil {
			return err
		}

		if len(updateUserName) == 0 {
			updateUserName = string(resp.Name)
		}
		if len(updateUserDescription) == 0 {
			updateUserDescription = resp.Description
		}
		if len(updateUserDefaultProject) == 0 && resp.DefaultProject != nil {
			updateUserDefaultProject = *resp.DefaultProject
		}

		var isUserEnabled bool

		if !enableUser && !disableUser {
			isUserEnabled = *resp.Enabled
		}
		if enableUser {
			isUserEnabled = true
		}
		if disableUser {
			isUserEnabled = false
		}

		resp, err = psOsClient.UpdateUser(ctx, args[0], openapi.UpdateOpenStackUser{
			Name:           &updateUserName,
			Description:    &updateUserDescription,
			DefaultProject: &updateUserDefaultProject,
			Password:       &updateUserPassword,
			Enabled:        &isUserEnabled,
		})

		if err != nil {
			return err
		}

		printUsers([]openapi.CreatedOpenStackUser{*resp})
		return nil
	},
}

func init() {
	userCmd.AddCommand(userUpdateCmd)

	userUpdateCmd.Flags().StringVarP(&updateUserName, "name", "n", "", "Specify the name of the user")
	userUpdateCmd.Flags().StringVarP(&updateUserDescription, "description", "d", "", "Specify the description of the user")
	userUpdateCmd.Flags().StringVar(&updateUserDefaultProject, "default-project", "", "Specify the default project of the user")
	userUpdateCmd.Flags().StringVarP(&updateUserPassword, "password", "p", "", "Specify the password of the user")
	//Unfortunately the API works via PUT -> Need to specify the password everytime
	if err := userUpdateCmd.MarkFlagRequired("password"); err != nil {
		panic(err)
	}

	userUpdateCmd.Flags().BoolVar(&enableUser, "enable", false, "Enable the specified user")
	userUpdateCmd.Flags().BoolVar(&disableUser, "disable", false, "Disable the specified user")
}
