/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Wieneo/pco-reseller-cli/v2/pkg/openapi"
	"github.com/spf13/cobra"
)

var enableProject, disableProject bool
var projectName string

// listCmd represents the list command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updated a projects name, description or enablement",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify the id of the project")
			os.Exit(1)
		}

		if len(args) > 1 {
			fmt.Println("Please only specify the id of the project")
			os.Exit(1)
		}

		if enableProject && disableProject {
			fmt.Println("Can't enable and disable the project")
			os.Exit(1)
		}

		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		resp, err := psOsClient.GetProject(ctx, args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
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
			fmt.Println(err.Error())
			os.Exit(1)
		}

		printProjects([]openapi.ProjectCreatedResponse{*resp})
	},
}

func init() {
	projectCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&projectName, "name", "n", "", "Update the name of the project")
	updateCmd.Flags().StringVarP(&projectDescription, "description", "d", "", "Update the description of the project")
	updateCmd.Flags().BoolVar(&enableProject, "enable", false, "Enable the specified project")
	updateCmd.Flags().BoolVar(&disableProject, "disable", false, "Disable the specified project")
}
