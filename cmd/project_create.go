/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Wieneo/pco-reseller-cli/v2/pkg/openapi"
	"github.com/spf13/cobra"
)

var projectDescription string
var projectWithDefaultNetwork bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new reseller project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify the name of the project")
			os.Exit(1)
		}

		if len(args) > 1 {
			fmt.Println("Please only specify the name of the project")
			os.Exit(1)
		}

		ctx := context.Background()

		enabled := true

		resp, err := psOsClient.CreateProject(ctx, openapi.ProjectCreate{
			Name:                args[0],
			Description:         projectDescription,
			Enabled:             &enabled,
			NetworkPreconfigure: &projectWithDefaultNetwork,
		})

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(resp.Id)
	},
}

func init() {
	projectCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&projectDescription, "description", "d", "No Description", "Specify the description of the project")

	createCmd.Flags().BoolVar(&projectWithDefaultNetwork, "with-default-network", false, "Specify if the default network should be created (default false)")
}
