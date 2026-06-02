package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
)

func printImages(images []openapi.ImageResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Status", "Visibility", "Disk Format", "Size"})

	for _, img := range images {
		name := ""
		if img.Name != nil {
			name = *img.Name
		}
		status := ""
		if img.Status != nil {
			status = *img.Status
		}
		visibility := ""
		if img.Visibility != nil {
			visibility = *img.Visibility
		}
		diskFormat := ""
		if img.DiskFormat != nil {
			diskFormat = *img.DiskFormat
		}
		size := 0
		if img.Size != nil {
			size = *img.Size
		}

		t.AppendRow([]interface{}{img.Id, name, status, visibility, diskFormat, size})
	}

	t.AppendFooter(table.Row{"", "", "", "Total", len(images)})
	t.Render()
}
