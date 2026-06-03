/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var projectUserAddCmd = &cobra.Command{
	Use:   "add [project-id] [user-id]",
	Short: "Add a user to a specific project",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		if err := psOsClient.AddUserToProject(ctx, args[0], args[1]); err != nil {
			return err
		}

		log.Infof("Added user %s to project %s", args[1], args[0])
		return nil
	},
}

func init() {
	projectUserCmd.AddCommand(projectUserAddCmd)
}
