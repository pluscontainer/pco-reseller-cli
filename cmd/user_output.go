package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/pluscontainer/reseller-cli/pkg/openapi"
)

func printUsers(users []openapi.CreatedOpenStackUser) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Description", "Default Project", "Enabled"})

	for _, k := range users {
		enabledString := "Enabled"
		if !*k.Enabled {
			enabledString = "Disabled"
		}

		t.AppendRow([]interface{}{k.Id, k.Name, k.Description, k.DefaultProject, enabledString})
	}

	t.AppendFooter(table.Row{"", "", "Total", len(users)})
	t.Render()
}
