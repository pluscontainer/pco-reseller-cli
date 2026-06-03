/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/sethvargo/go-password/password"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var projectDescription string
var projectWithDefaultNetwork bool
var enableProject, disableProject bool
var projectName string
var bootstrapUserName, bootstrapPassword, bootstrapDescription string
var bootstrapWithDefaultNetwork, bootstrapWithDefaultQuota bool

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
}

var createCmd = &cobra.Command{
	Use:     "create [project-name]",
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

		printProjects([]openapi.ProjectCreatedResponse{*resp})
		return nil
	},
}

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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetProjects(ctx)
		if err != nil {
			return err
		}

		printProjects(*resp)
		return nil
	},
}

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

		if len(projectName) == 0 {
			// domain gets appended to project name — omit when sending the PUT request
			projectName = strings.Join(strings.Split(resp.Name, "-")[1:], "-")
		}
		if len(projectDescription) == 0 {
			projectDescription = resp.Description
		}

		var isEnabled bool
		if !enableProject && !disableProject {
			isEnabled = *resp.Enabled
		} else {
			isEnabled = enableProject
		}

		resp, err = psOsClient.UpdateProject(ctx, args[0], openapi.ProjectUpdate{
			Name:        &projectName,
			Description: &projectDescription,
			Enabled:     &isEnabled,
		})
		if err != nil {
			return err
		}

		printProjects([]openapi.ProjectCreatedResponse{*resp})
		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [project-id]",
	Short: "Delete a reseller project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		if err := psOsClient.DeleteProject(ctx, args[0]); err != nil {
			return err
		}

		log.Infof("Deleted project %s", args[0])
		return nil
	},
}

var bootstrapCmd = &cobra.Command{
	Use:     "bootstrap [project-name]",
	Short:   "Bootstrap a project with a user",
	Long:    `Creates a project, a user and assigns the user to the project in one step.`,
	Example: "  pco-reseller-cli project bootstrap my-project\n  pco-reseller-cli project bootstrap my-project --password secret\n  pco-reseller-cli project bootstrap my-project --user-name my-user --with-default-network",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		description := bootstrapDescription
		if !cmd.Flags().Changed("description") {
			description = projectName
		}

		userName := bootstrapUserName
		if !cmd.Flags().Changed("user-name") {
			userName = projectName + "-admin"
		}

		pw := bootstrapPassword
		if !cmd.Flags().Changed("password") {
			var err error
			pw, err = password.Generate(24, 4, 4, false, true)
			if err != nil {
				return fmt.Errorf("error generating password: %w", err)
			}
		}

		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		enabled := true

		project, err := psOsClient.CreateProject(ctx, openapi.ProjectCreate{
			Name:                projectName,
			Description:         description,
			Enabled:             &enabled,
			NetworkPreconfigure: &bootstrapWithDefaultNetwork,
		})
		if err != nil {
			return fmt.Errorf("error creating project: %w", err)
		}
		log.Infof("Created project %s (%s)", project.Name, project.Id)

		user, err := psOsClient.CreateUser(ctx, openapi.CreateOpenStackUser{
			Name:           userName,
			Description:    description,
			Enabled:        &enabled,
			DefaultProject: &project.Id,
			Password:       pw,
		})
		if err != nil {
			return fmt.Errorf("error creating user: %w", err)
		}
		log.Infof("Created user %s (%s)", user.Name, user.Id)

		if err := psOsClient.AddUserToProject(ctx, project.Id, user.Id); err != nil {
			return fmt.Errorf("error adding user to project: %w", err)
		}
		log.Infof("Assigned user %s to project %s", user.Name, project.Name)

		if bootstrapWithDefaultQuota {
			if _, err := psOsClient.UpdateProjectQuota(ctx, project.Id, defaultQuota); err != nil {
				return fmt.Errorf("error applying default quota: %w", err)
			}
			log.Infof("Applied default quota to project %s", project.Name)
		}

		fmt.Println()
		fmt.Println("Bootstrap completed successfully")
		fmt.Println("--------------------------------")
		fmt.Printf("Project Name: %s\n", project.Name)
		fmt.Printf("Project ID:   %s\n", project.Id)
		fmt.Printf("User Name:    %s\n", user.Name)
		fmt.Printf("Password:     %s\n", pw)
		return nil
	},
}

func printProjects(projects []openapi.ProjectCreatedResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Description", "Enabled"})

	for _, k := range projects {
		enabledString := "Enabled"
		if !*k.Enabled {
			enabledString = "Disabled"
		}
		t.AppendRow([]any{k.Id, k.Name, k.Description, enabledString})
	}

	t.AppendFooter(table.Row{"", "", "Total", len(projects)})
	t.Render()
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(createCmd, getCmd, listCmd, updateCmd, deleteCmd, bootstrapCmd)

	createCmd.Flags().StringVarP(&projectDescription, "description", "d", "", "Description of the project (defaults to the project name if not set)")
	createCmd.Flags().BoolVar(&projectWithDefaultNetwork, "with-default-network", false, "Preconfigure the project with a default network, router and security groups")

	updateCmd.Flags().StringVarP(&projectName, "name", "n", "", "Update the name of the project")
	updateCmd.Flags().StringVarP(&projectDescription, "description", "d", "", "Update the description of the project")
	updateCmd.Flags().BoolVar(&enableProject, "enable", false, "Enable the specified project")
	updateCmd.Flags().BoolVar(&disableProject, "disable", false, "Disable the specified project")
	updateCmd.MarkFlagsMutuallyExclusive("enable", "disable")

	bootstrapCmd.Flags().StringVarP(&bootstrapDescription, "description", "d", "", "Description of the project and user (defaults to the project name if not set)")
	bootstrapCmd.Flags().StringVar(&bootstrapUserName, "user-name", "", "Name of the user to create (defaults to <project-name>-admin)")
	bootstrapCmd.Flags().StringVarP(&bootstrapPassword, "password", "p", "", "Password for the new user (auto-generated if not set)")
	bootstrapCmd.Flags().BoolVar(&bootstrapWithDefaultNetwork, "with-default-network", false, "Preconfigure the project with a default network, router and security groups")
	bootstrapCmd.Flags().BoolVar(&bootstrapWithDefaultQuota, "with-default-quota", false, "Apply the standard quota set to the project")
}
