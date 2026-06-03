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

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Manage images",
}

func printImages(images []openapi.ImageResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Status", "Visibility", "Disk Format", "Size"})

	for _, img := range images {
		name, status, visibility, diskFormat := "", "", "", ""
		size := 0
		if img.Name != nil {
			name = *img.Name
		}
		if img.Status != nil {
			status = *img.Status
		}
		if img.Visibility != nil {
			visibility = *img.Visibility
		}
		if img.DiskFormat != nil {
			diskFormat = *img.DiskFormat
		}
		if img.Size != nil {
			size = *img.Size
		}
		t.AppendRow([]any{img.Id, name, status, visibility, diskFormat, size})
	}

	t.AppendFooter(table.Row{"", "", "", "Total", len(images)})
	t.Render()
}

func init() {
	rootCmd.AddCommand(imageCmd)
}
