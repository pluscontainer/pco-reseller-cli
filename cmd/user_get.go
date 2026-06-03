/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var userGetCmd = &cobra.Command{
	Use:   "get [user-id]",
	Short: "Get a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetUser(ctx, args[0])
		if err != nil {
			return err
		}

		printUsers([]openapi.CreatedOpenStackUser{*resp})
		return nil
	},
}

func init() {
	userCmd.AddCommand(userGetCmd)
}
