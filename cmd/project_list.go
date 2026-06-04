/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetProjects(ctx)
		if err != nil {
			return err
		}

		printProjects(*resp)
		return nil
	},
}

func init() {
	projectCmd.AddCommand(listCmd)
}
