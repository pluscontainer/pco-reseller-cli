/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var imageVisibility string

var imageUpdateVisibilityCmd = &cobra.Command{
	Use:   "update-visibility",
	Short: "Update image visibility",
	Long:  `Update the visibility of an image (community, public, private, shared)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}

		if len(args) > 1 {
			fmt.Fprintln(os.Stderr, "Error: too many arguments, expected exactly one image ID")
			os.Exit(1)
		}

		visibility := openapi.ImageVisibility(imageVisibility)
		if !visibility.Valid() {
			fmt.Printf("Invalid visibility %q. Valid values: community, public, private, shared\n", imageVisibility)
			os.Exit(1)
		}

		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		if err := psOsClient.UpdateImageVisibility(ctx, args[0], visibility); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		resp, err := psOsClient.GetImage(ctx, args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		printImages([]openapi.ImageResponse{*resp})
	},
}

func init() {
	imageCmd.AddCommand(imageUpdateVisibilityCmd)

	imageUpdateVisibilityCmd.Flags().StringVarP(&imageVisibility, "visibility", "v", "", "Visibility of the image (community, public, private, shared)")
	imageUpdateVisibilityCmd.MarkFlagRequired("visibility")
}
