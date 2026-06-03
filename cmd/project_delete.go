/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var deleteCmd = &cobra.Command{
	Use:   "delete [project-id]",
	Short: "Delete a reseller project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		err := psOsClient.DeleteProject(ctx, args[0])

		if err != nil {
			return err
		}

		fmt.Println("Project deleted")
		return nil
	},
}

func init() {
	projectCmd.AddCommand(deleteCmd)
}
