/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var applyDefaultQuotaCmd = &cobra.Command{
	Use:   "apply-default [project-id]",
	Short: "Apply the standard quota set to the specified project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.UpdateProjectQuota(ctx, args[0], defaultQuota)
		if err != nil {
			return err
		}

		printQuota(*resp)
		return nil
	},
}

func init() {
	quotaCmd.AddCommand(applyDefaultQuotaCmd)
}
