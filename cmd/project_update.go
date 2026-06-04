/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"strings"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var updateProjectName, updateProjectDescription string
var enableProject, disableProject bool

var updateCmd = &cobra.Command{
	Use:   "update [project-id]",
	Short: "Update a project's name, description or enablement",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetProject(ctx, args[0])
		if err != nil {
			return err
		}

		if len(updateProjectName) == 0 {
			// domain gets appended to project name — omit when sending the PUT request
			updateProjectName = strings.Join(strings.Split(resp.Name, "-")[1:], "-")
		}
		if len(updateProjectDescription) == 0 {
			updateProjectDescription = resp.Description
		}

		var isEnabled bool
		if !enableProject && !disableProject {
			isEnabled = *resp.Enabled
		} else {
			isEnabled = enableProject
		}

		resp, err = psOsClient.UpdateProject(ctx, args[0], openapi.ProjectUpdate{
			Name:        &updateProjectName,
			Description: &updateProjectDescription,
			Enabled:     &isEnabled,
		})
		if err != nil {
			return err
		}

		printProjects([]openapi.ProjectCreatedResponse{*resp})
		return nil
	},
}

func init() {
	projectCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&updateProjectName, "name", "n", "", "Update the name of the project")
	updateCmd.Flags().StringVarP(&updateProjectDescription, "description", "d", "", "Update the description of the project")
	updateCmd.Flags().BoolVar(&enableProject, "enable", false, "Enable the specified project")
	updateCmd.Flags().BoolVar(&disableProject, "disable", false, "Disable the specified project")
	updateCmd.MarkFlagsMutuallyExclusive("enable", "disable")
}
