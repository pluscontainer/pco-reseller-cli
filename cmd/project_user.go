/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var projectUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users accessing projects",
	Long:  `List, assign or revoke access to OpenStack projects`,
}

func init() {
	projectCmd.AddCommand(projectUserCmd)
}
