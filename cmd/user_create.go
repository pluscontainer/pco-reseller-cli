/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/pluscloudopen/reseller-cli/v2/pkg/openapi"
	"github.com/spf13/cobra"
)

var userDescription, userDefaultProject, userPassword string

// createCmd represents the create command
var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new reseller user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify the name of the user")
			os.Exit(1)
		}

		if len(args) > 1 {
			fmt.Println("Please only specify the name of the user")
			os.Exit(1)
		}

		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		enabled := true

		resp, err := psOsClient.CreateUser(ctx, openapi.CreateOpenStackUser{
			Name:           types.Email(args[0]),
			Description:    userDescription,
			Enabled:        &enabled,
			DefaultProject: &userDefaultProject,
			Password:       userPassword,
		})

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(resp.Id)
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)

	userCreateCmd.Flags().StringVarP(&userDescription, "description", "d", "No Description", "Specify the description of the user")

	userCreateCmd.Flags().StringVar(&userDefaultProject, "default-project", "", "Specify the default project of the user")
	userCreateCmd.MarkFlagRequired("default-project")

	userCreateCmd.Flags().StringVarP(&userPassword, "password", "p", "", "Specify the password of the user")
	userCreateCmd.MarkFlagRequired("password")
}
