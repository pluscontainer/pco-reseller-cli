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
var userDeleteCmd = &cobra.Command{
	Use:   "delete [user-id]",
	Short: "Delete a reseller user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()

		err := psOsClient.DeleteUser(ctx, args[0])

		if err != nil {
			return err
		}

		fmt.Printf("Deleted user %s\n", args[0])
		return nil
	},
}

func init() {
	userCmd.AddCommand(userDeleteCmd)
}
