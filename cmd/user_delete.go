/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var userDeleteCmd = &cobra.Command{
	Use:   "delete [user-id]",
	Short: "Delete a reseller user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		if err := psOsClient.DeleteUser(ctx, args[0]); err != nil {
			return err
		}

		log.Infof("Deleted user %s", args[0])
		return nil
	},
}

func init() {
	userCmd.AddCommand(userDeleteCmd)
}
