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

var enableProject, disableProject bool
var projectName string

// listCmd represents the list command
var updateCmd = &cobra.Command{
	Use:   "update [project-id]",
	Short: "Updated a projects name, description or enablement",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		resp, err := psOsClient.GetProject(ctx, args[0])
		if err != nil {
			return err
		}

		var isProjectEnabled bool

		if len(projectName) == 0 {
			//domain gets appended to project name -> Need to omit that when sending the put request
			projectName = strings.Join(strings.Split(resp.Name, "-")[1:], "-")
		}
		if len(projectDescription) == 0 {
			projectDescription = resp.Description
		}

		if !enableProject && !disableProject {
			isProjectEnabled = *resp.Enabled
		}
		if enableProject {
			isProjectEnabled = true
		}
		if disableProject {
			isProjectEnabled = false
		}

		resp, err = psOsClient.UpdateProject(ctx, args[0], openapi.ProjectUpdate{
			Name:        &projectName,
			Description: &projectDescription,
			Enabled:     &isProjectEnabled,
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

	updateCmd.Flags().StringVarP(&projectName, "name", "n", "", "Update the name of the project")
	updateCmd.Flags().StringVarP(&projectDescription, "description", "d", "", "Update the description of the project")
	updateCmd.Flags().BoolVar(&enableProject, "enable", false, "Enable the specified project")
	updateCmd.Flags().BoolVar(&disableProject, "disable", false, "Disable the specified project")
	updateCmd.MarkFlagsMutuallyExclusive("enable", "disable")
}
