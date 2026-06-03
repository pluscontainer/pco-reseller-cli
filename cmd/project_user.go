/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var projectUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users accessing projects",
	Long:  `List, assign or revoke access to OpenStack projects`,
}

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

		log.Infof("Removed user %s from project %s", args[1], args[0])
		return nil
	},
}

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
			log.Info("No users assigned to project")
			return nil
		}

		for _, k := range *resp {
			fmt.Println(k.User)
		}
		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectUserCmd)
	projectUserCmd.AddCommand(projectUserAddCmd, projectUserDeleteCmd, projectUserListCmd)
}
