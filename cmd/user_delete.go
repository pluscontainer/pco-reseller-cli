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

// createCmd represents the create command
var userDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a reseller user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify the id of the user")
			os.Exit(1)
		}

		if len(args) > 1 {
			fmt.Println("Please only specify the id of the user")
			os.Exit(1)
		}

		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()

		err := psOsClient.DeleteUser(ctx, args[0])

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("User deleted")
	},
}

func init() {
	userCmd.AddCommand(userDeleteCmd)
}
