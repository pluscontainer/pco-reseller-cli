/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [project-id]",
	Short: "Get a project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetProject(ctx, args[0])
		if err != nil {
			return err
		}

		printProjects([]openapi.ProjectCreatedResponse{*resp})
		return nil
	},
}

func init() {
	projectCmd.AddCommand(getCmd)
}
