/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var projectUserDeleteCmd = &cobra.Command{
	Use:   "delete [project-id] [user-id]",
	Short: "Remove a user from a specific project",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		if err := psOsClient.RemoveUserFromProject(ctx, args[0], args[1]); err != nil {
			return err
		}

		fmt.Printf("Removed user %s from project %s\n", args[1], args[0])
		return nil
	},
}

func init() {
	projectUserCmd.AddCommand(projectUserDeleteCmd)
}
