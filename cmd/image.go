/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var imageVisibility string

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Manage images",
}

var imageGetCmd = &cobra.Command{
	Use:   "get [image-id]",
	Short: "Get an image by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetImage(ctx, args[0])
		if err != nil {
			return err
		}

		printImages([]openapi.ImageResponse{*resp})
		return nil
	},
}

var imageUpdateVisibilityCmd = &cobra.Command{
	Use:   "update-visibility [image-id]",
	Short: "Update image visibility",
	Long:  `Update the visibility of an image (community, public, private, shared)`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		visibility := openapi.ImageVisibility(imageVisibility)
		if !visibility.Valid() {
			return fmt.Errorf("invalid visibility %q — valid values: community, public, private, shared", imageVisibility)
		}

		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		if err := psOsClient.UpdateImageVisibility(ctx, args[0], visibility); err != nil {
			return err
		}

		resp, err := psOsClient.GetImage(ctx, args[0])
		if err != nil {
			return err
		}

		printImages([]openapi.ImageResponse{*resp})
		return nil
	},
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
	imageCmd.AddCommand(imageGetCmd, imageUpdateVisibilityCmd)

	imageUpdateVisibilityCmd.Flags().StringVarP(&imageVisibility, "visibility", "v", "", "Visibility of the image (community, public, private, shared)")
	imageUpdateVisibilityCmd.MarkFlagRequired("visibility")
}
