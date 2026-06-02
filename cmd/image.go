/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Manage images",
	Long:  `Get images and manage their visibility`,
}

func init() {
	rootCmd.AddCommand(imageCmd)
}
