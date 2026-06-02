/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var projectDescription string
var projectWithDefaultNetwork bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new reseller project",
	Example: "  pco-reseller-cli project create my-project\n  pco-reseller-cli project create my-project --description \"my project\" --with-default-network",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}

		if len(args) > 1 {
			fmt.Fprintln(os.Stderr, "Error: too many arguments, expected exactly one project name")
			os.Exit(1)
		}

		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		enabled := true

		description := projectDescription
		if !cmd.Flags().Changed("description") {
			description = args[0]
		}

		resp, err := psOsClient.CreateProject(ctx, openapi.ProjectCreate{
			Name:                args[0],
			Description:         description,
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

	createCmd.Flags().StringVarP(&projectDescription, "description", "d", "", "Description of the project (defaults to the project name if not set)")

	createCmd.Flags().BoolVar(&projectWithDefaultNetwork, "with-default-network", false, "Specify if the default network should be created (default false)")
}
