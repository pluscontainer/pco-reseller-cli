/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/plusserver/pluscloudopen-reseller-cli/v2/pkg/openapi"
	"github.com/spf13/cobra"
)

var updateUserName, updateUserDescription, updateUserDefaultProject, updateUserPassword string

var enableUser, disableUser bool

// createCmd represents the create command
var userUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a reseller user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify the id of the user")
			os.Exit(1)
		}

		if len(args) > 1 {
			fmt.Println("Please only specify the id of the user")
			os.Exit(1)
		}

		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()

		resp, err := psOsClient.GetUser(ctx, args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
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

		email := types.Email(updateUserName)

		resp, err = psOsClient.UpdateUser(ctx, args[0], openapi.UpdateOpenStackUser{
			Name:           &email,
			Description:    &updateUserDescription,
			DefaultProject: &updateUserDefaultProject,
			Password:       &updateUserPassword,
			Enabled:        &isUserEnabled,
		})

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		printUsers([]openapi.CreatedOpenStackUser{*resp})
	},
}

func init() {
	userCmd.AddCommand(userUpdateCmd)

	userUpdateCmd.Flags().StringVarP(&updateUserName, "name", "n", "", "Specify the name of the user")
	userUpdateCmd.Flags().StringVarP(&updateUserDescription, "description", "d", "", "Specify the description of the user")
	userUpdateCmd.Flags().StringVar(&updateUserDefaultProject, "default-project", "", "Specify the default project of the user")
	userUpdateCmd.Flags().StringVarP(&updateUserPassword, "password", "p", "", "Specify the password of the user")
	//Unfortunately the API works via PUT -> Need to specify the password everytime
	userUpdateCmd.MarkFlagRequired("password")

	userUpdateCmd.Flags().BoolVar(&enableUser, "enable", false, "Enable the specified user")
	userUpdateCmd.Flags().BoolVar(&disableUser, "disable", false, "Disable the specified user")
}
