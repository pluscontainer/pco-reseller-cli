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

// listCmd represents the list command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}

		if len(args) > 1 {
			fmt.Fprintln(os.Stderr, "Error: too many arguments, expected exactly one project ID")
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
