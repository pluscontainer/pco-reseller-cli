/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check API Availability",
	Long: `Check if the current configuration is correct and the corresponding API is available.
Exits with code != 0 if an error occurs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fetchPsOpenStackClientOrDie()

		//If we get here -> Login has succeded
		fmt.Println("Ping successful.")
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
