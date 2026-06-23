/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var createProjectDescription string
var createProjectWithDefaultNetwork bool

var createCmd = &cobra.Command{
	Use:     "create [project-name]",
	Short:   "Create a new reseller project",
	Example: "  pco-reseller-cli project create my-project\n  pco-reseller-cli project create my-project --description \"my project\" --with-default-network",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		enabled := true

		description := createProjectDescription
		if !cmd.Flags().Changed("description") {
			description = args[0]
		}

		resp, err := psOsClient.CreateProject(ctx, openapi.ProjectCreate{
			Name:                args[0],
			Description:         description,
			Enabled:             &enabled,
			NetworkPreconfigure: &createProjectWithDefaultNetwork,
		})
		if err != nil {
			return err
		}

		printProjects([]openapi.ProjectCreatedResponse{*resp})
		return nil
	},
}

func init() {
	projectCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&createProjectDescription, "description", "d", "", "Description of the project (defaults to the project name if not set)")
	createCmd.Flags().BoolVar(&createProjectWithDefaultNetwork, "with-default-network", false, "Preconfigure the project with a default network, router and security groups")
}
