/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
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
}
