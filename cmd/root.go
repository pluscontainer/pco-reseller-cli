/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/pluscontainer/pco-reseller-cli/pkg/psos"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pco-reseller-cli",
	Short: "Managed OpenStack projects and users as a reseller",
	Long:  `Managed OpenStack projects and users as a reseller`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const (
	psosEndpointEnvKey = "PSOS_ENDPOINT"
	psosUsernameEnvKey = "PSOS_USERNAME"
	psosPasswordEnvKey = "PSOS_PASSWORD"
)

func fetchPsOpenStackClientOrDie() *psos.PsOpenstackClient {
	errList := []error{}
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

	var err error
	psOsClient, err := psos.Login(endpoint, username, password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return psOsClient
}

func envKeyMissingError(key string) error {
	return fmt.Errorf("please define env %s", key)
}
