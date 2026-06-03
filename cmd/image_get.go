/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

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

func init() {
	imageCmd.AddCommand(imageGetCmd)
}
