/*
Copyright © 2022 PlusServer GmbH

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long:  `Create, update, delete projects`,
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
