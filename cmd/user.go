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

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long:  `Get, create, update, delete users`,
}

func printUsers(users []openapi.CreatedOpenStackUser) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Description", "Default Project", "Enabled"})

	for _, k := range users {
		enabledString := "Enabled"
		if !*k.Enabled {
			enabledString = "Disabled"
		}
		t.AppendRow([]any{k.Id, k.Name, k.Description, k.DefaultProject, enabledString})
	}

	t.AppendFooter(table.Row{"", "", "Total", len(users)})
	t.Render()
}

func init() {
	rootCmd.AddCommand(userCmd)
}
