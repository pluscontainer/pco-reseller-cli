package cmd

import (
	"os"

	"github.com/Wieneo/pco-reseller-cli/v2/pkg/openapi"
	"github.com/jedib0t/go-pretty/table"
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
