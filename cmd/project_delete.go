/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a reseller project",
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
		err := psOsClient.DeleteProject(ctx, args[0])

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Project deleted")
	},
}

func init() {
	projectCmd.AddCommand(deleteCmd)
}
