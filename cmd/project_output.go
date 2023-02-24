package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/pluscloudopen/reseller-cli/v2/pkg/openapi"
)

func printProjects(projects []openapi.ProjectCreatedResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Description", "Enabled"})

	for _, k := range projects {
		enabledString := "Enabled"
		if !*k.Enabled {
			enabledString = "Disabled"
		}

		t.AppendRow([]interface{}{k.Id, k.Name, k.Description, enabledString})
	}

	t.AppendFooter(table.Row{"", "", "Total", len(projects)})
	t.Render()
}
