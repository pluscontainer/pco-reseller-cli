/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var getQuotaCmd = &cobra.Command{
	Use:   "get [project-id]",
	Short: "Get the quotas of the specified project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetProjectQuota(ctx, args[0])
		if err != nil {
			return err
		}

		printQuota(*resp)
		return nil
	},
}

func init() {
	quotaCmd.AddCommand(getQuotaCmd)
}
