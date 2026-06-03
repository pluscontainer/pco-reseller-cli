/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [project-id]",
	Short: "Delete a reseller project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		if err := psOsClient.DeleteProject(ctx, args[0]); err != nil {
			return err
		}

		log.Infof("Deleted project %s", args[0])
		return nil
	},
}

func init() {
	projectCmd.AddCommand(deleteCmd)
}
