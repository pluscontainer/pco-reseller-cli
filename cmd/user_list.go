/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users assigned to the reseller account",
	Long:  `List all users assigned to the reseller account`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
