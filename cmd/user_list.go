/*
Copyright © 2022 PlusServer GmbH

*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users assigned to the reseller account",
	Long:  `List all users assigned to the reseller account`,
	Run: func(cmd *cobra.Command, args []string) {
		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		resp, err := psOsClient.GetUsers(ctx)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		printUsers(*resp)
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
