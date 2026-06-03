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

		var isEnabled bool
		if !enableUser && !disableUser {
			isEnabled = *resp.Enabled
		} else {
			isEnabled = enableUser
		}

		resp, err = psOsClient.UpdateUser(ctx, args[0], openapi.UpdateOpenStackUser{
			Name:           &updateUserName,
			Description:    &updateUserDescription,
			DefaultProject: &updateUserDefaultProject,
			Password:       &updateUserPassword,
			Enabled:        &isEnabled,
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
	userUpdateCmd.Flags().StringVarP(&updateUserName, "name", "n", "", "Update the name of the user")
	userUpdateCmd.Flags().StringVarP(&updateUserDescription, "description", "d", "", "Update the description of the user")
	userUpdateCmd.Flags().StringVar(&updateUserDefaultProject, "default-project", "", "Update the default project of the user")
	userUpdateCmd.Flags().StringVarP(&updateUserPassword, "password", "p", "", "Password of the user")
	if err := userUpdateCmd.MarkFlagRequired("password"); err != nil {
		panic(err)
	}
	userUpdateCmd.Flags().BoolVar(&enableUser, "enable", false, "Enable the specified user")
	userUpdateCmd.Flags().BoolVar(&disableUser, "disable", false, "Disable the specified user")
	userUpdateCmd.MarkFlagsMutuallyExclusive("enable", "disable")
}
