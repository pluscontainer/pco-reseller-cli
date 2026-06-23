/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users assigned to the reseller account",
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetUsers(ctx)
		if err != nil {
			return err
		}

		printUsers(*resp)
		return nil
	},
}

func init() {
	userCmd.AddCommand(userListCmd)
}
