/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var projectUserAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a user to a specific project",
	Long:  `Add a user to a specific project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Please specify the id of the project and the user")
			os.Exit(1)
		}

		psOsClient := fetchPsOpenStackClientOrDie()

		ctx := context.Background()
		err := psOsClient.AddUserToProject(ctx, args[0], args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("User added to project")
	},
}

func init() {
	userCmd.AddCommand(projectUserAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
