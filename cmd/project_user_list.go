/*
Copyright © 2022 PlusServer GmbH

*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var projectUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users assigned to the specified project",
	Long:  `List all users assigned to the specified project`,
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
		resp, err := psOsClient.GetUsersInProject(ctx, args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if resp == nil {
			fmt.Println("No users assigned to project")
			return
		}

		for _, k := range *resp {
			fmt.Println(k.User)
		}
	},
}

func init() {
	projectUserCmd.AddCommand(projectUserListCmd)
}
