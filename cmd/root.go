/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pluscontainer/pco-reseller-cli/pkg/psos"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type cliFormatter struct{}

func (f *cliFormatter) Format(entry *log.Entry) ([]byte, error) {
	level := strings.ToUpper(entry.Level.String())
	return []byte(fmt.Sprintf("[%s] %s\n", level, entry.Message)), nil
}

func init() {
	log.SetFormatter(&cliFormatter{})
}

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

func fetchPsOpenStackClientOrDie() *psos.PsOpenstackClient {
	var missing []string
	for _, key := range []string{psosEndpointEnvKey, psosUsernameEnvKey, psosPasswordEnvKey} {
		if _, ok := os.LookupEnv(key); !ok {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		for _, key := range missing {
			log.Errorf("environment variable %s is not set", key)
		}
		os.Exit(1)
	}

	psOsClient, err := psos.Login(
		os.Getenv(psosEndpointEnvKey),
		os.Getenv(psosUsernameEnvKey),
		os.Getenv(psosPasswordEnvKey),
	)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	return psOsClient
}
