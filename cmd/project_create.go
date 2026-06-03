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

var projectDescription string
var projectWithDefaultNetwork bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create [project-id]",
	Short:   "Create a new reseller project",
	Example: "  pco-reseller-cli project create my-project\n  pco-reseller-cli project create my-project --description \"my project\" --with-default-network",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
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
			return err
		}

		fmt.Println(resp.Id)
		return nil
	},
}

func init() {
	projectCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&projectDescription, "description", "d", "", "Description of the project (defaults to the project name if not set)")

	createCmd.Flags().BoolVar(&projectWithDefaultNetwork, "with-default-network", false, "Specify if the default network should be created (default false)")
}
