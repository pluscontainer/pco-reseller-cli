/*
Copyright © 2022 PlusServer GmbH

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var projectUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users accessing projects",
	Long:  `List, assign or revoke access to OpenStack projects`,
}

func init() {
	projectCmd.AddCommand(projectUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
