/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var projectUserAddCmd = &cobra.Command{
	Use:   "add [project-id] [user-id]",
	Short: "Add a user to a specific project",
	Long:  `Add a user to a specific project`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		err := psOsClient.AddUserToProject(ctx, args[0], args[1])
		if err != nil {
			return err
		}

		fmt.Printf("Added user %s to project %s\n", args[1], args[0])
		return nil
	},
}

func init() {
	projectUserCmd.AddCommand(projectUserAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
