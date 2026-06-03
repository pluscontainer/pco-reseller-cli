/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var imageVisibility string

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

func init() {
	imageCmd.AddCommand(imageUpdateVisibilityCmd)
	imageUpdateVisibilityCmd.Flags().StringVarP(&imageVisibility, "visibility", "v", "", "Visibility of the image (community, public, private, shared)")
	imageUpdateVisibilityCmd.MarkFlagRequired("visibility")
}
