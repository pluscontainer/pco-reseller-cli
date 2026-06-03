/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/pluscontainer/pco-reseller-cli/pkg/psos"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "pco-reseller-cli",
	Short:         "Manage OpenStack projects and users as a reseller",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() error {
	return rootCmd.Execute()
}

const (
	psosEndpointEnvKey = "PSOS_ENDPOINT"
	psosUsernameEnvKey = "PSOS_USERNAME"
	psosPasswordEnvKey = "PSOS_PASSWORD"
)

func envKeyMissingError(key string) error {
	return fmt.Errorf("environment variable %s is not set", key)
}

func fetchPsOpenStackClientOrDie() *psos.PsOpenstackClient {
	var errList []error
	endpoint, ok := os.LookupEnv(psosEndpointEnvKey)
	if !ok {
		errList = append(errList, envKeyMissingError(psosEndpointEnvKey))
	}

	username, ok := os.LookupEnv(psosUsernameEnvKey)
	if !ok {
		errList = append(errList, envKeyMissingError(psosUsernameEnvKey))
	}

	password, ok := os.LookupEnv(psosPasswordEnvKey)
	if !ok {
		errList = append(errList, envKeyMissingError(psosPasswordEnvKey))
	}

	if err := errors.Join(errList...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	psOsClient, err := psos.Login(endpoint, username, password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return psOsClient
}
