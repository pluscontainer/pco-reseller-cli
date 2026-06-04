/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check API availability",
	Long:  `Check if the current configuration is correct and the API is reachable. Exits with a non-zero code on failure.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fetchPsOpenStackClientOrDie()
		fmt.Println("Ping successful.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
