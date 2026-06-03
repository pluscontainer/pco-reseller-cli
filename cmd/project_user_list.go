/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var projectUserListCmd = &cobra.Command{
	Use:   "list [project-id]",
	Short: "List all users assigned to the specified project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetUsersInProject(ctx, args[0])
		if err != nil {
			return err
		}

		if resp == nil {
			fmt.Println("No users assigned to project")
			return nil
		}

		for _, k := range *resp {
			fmt.Println(k.User)
		}
		return nil
	},
}

func init() {
	projectUserCmd.AddCommand(projectUserListCmd)
}
