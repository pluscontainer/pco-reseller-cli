/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/pluscloudopen/reseller-cli/v2/pkg/openapi"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify the id of the project")
			os.Exit(1)
		}

		if len(args) > 1 {
			fmt.Println("Please only specify the id of the project")
			os.Exit(1)
		}

		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		resp, err := psOsClient.GetProject(ctx, args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		printProjects([]openapi.ProjectCreatedResponse{*resp})
	},
}

func init() {
	projectCmd.AddCommand(getCmd)
}
