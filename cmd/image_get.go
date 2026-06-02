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

var imageGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an image",
	Long:  `Get an image by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify the id of the image")
			os.Exit(1)
		}

		if len(args) > 1 {
			fmt.Println("Please only specify the id of the image")
			os.Exit(1)
		}

		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		resp, err := psOsClient.GetImage(ctx, args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		printImages([]openapi.ImageResponse{*resp})
	},
}

func init() {
	imageCmd.AddCommand(imageGetCmd)
}
