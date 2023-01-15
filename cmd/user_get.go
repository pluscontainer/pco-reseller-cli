/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/plusserver/pluscloudopen-reseller-cli/v2/pkg/openapi"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a user",
	Long:  `Get a user`,
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
		resp, err := psOsClient.GetUser(ctx, args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		printUsers([]openapi.CreatedOpenStackUser{*resp})
	},
}

func init() {
	userCmd.AddCommand(userGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
