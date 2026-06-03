/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var userDescription, userDefaultProject, userPassword string
var updateUserName, updateUserDescription, updateUserDefaultProject, updateUserPassword string
var enableUser, disableUser bool

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long:  `Get, create, update, delete users`,
}

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

var userGetCmd = &cobra.Command{
	Use:   "get [user-id]",
	Short: "Get a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetUser(ctx, args[0])
		if err != nil {
			return err
		}

		printUsers([]openapi.CreatedOpenStackUser{*resp})
		return nil
	},
}

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users assigned to the reseller account",
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetUsers(ctx)
		if err != nil {
			return err
		}

		printUsers(*resp)
		return nil
	},
}

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

var userDeleteCmd = &cobra.Command{
	Use:   "delete [user-id]",
	Short: "Delete a reseller user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		if err := psOsClient.DeleteUser(ctx, args[0]); err != nil {
			return err
		}

		log.Infof("Deleted user %s", args[0])
		return nil
	},
}

func printUsers(users []openapi.CreatedOpenStackUser) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Description", "Default Project", "Enabled"})

	for _, k := range users {
		enabledString := "Enabled"
		if !*k.Enabled {
			enabledString = "Disabled"
		}
		t.AppendRow([]any{k.Id, k.Name, k.Description, k.DefaultProject, enabledString})
	}

	t.AppendFooter(table.Row{"", "", "Total", len(users)})
	t.Render()
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userCreateCmd, userGetCmd, userListCmd, userUpdateCmd, userDeleteCmd)

	userCreateCmd.Flags().StringVarP(&userDescription, "description", "d", "", "Description of the user (defaults to the user name if not set)")
	userCreateCmd.Flags().StringVar(&userDefaultProject, "default-project", "", "Default project of the user")
	if err := userCreateCmd.MarkFlagRequired("default-project"); err != nil {
		panic(err)
	}
	userCreateCmd.Flags().StringVarP(&userPassword, "password", "p", "", "Password of the user")
	if err := userCreateCmd.MarkFlagRequired("password"); err != nil {
		panic(err)
	}

	userUpdateCmd.Flags().StringVarP(&updateUserName, "name", "n", "", "Update the name of the user")
	userUpdateCmd.Flags().StringVarP(&updateUserDescription, "description", "d", "", "Update the description of the user")
	userUpdateCmd.Flags().StringVar(&updateUserDefaultProject, "default-project", "", "Update the default project of the user")
	userUpdateCmd.Flags().StringVarP(&updateUserPassword, "password", "p", "", "Password of the user")
	// API works via PUT — password must always be specified
	if err := userUpdateCmd.MarkFlagRequired("password"); err != nil {
		panic(err)
	}
	userUpdateCmd.Flags().BoolVar(&enableUser, "enable", false, "Enable the specified user")
	userUpdateCmd.Flags().BoolVar(&disableUser, "disable", false, "Disable the specified user")
	userUpdateCmd.MarkFlagsMutuallyExclusive("enable", "disable")
}
