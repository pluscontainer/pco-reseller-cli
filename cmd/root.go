/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Wieneo/pco-reseller-cli/v2/pkg/psos"
	"github.com/spf13/cobra"
)

var psOsClient *psos.PsOpenstackClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "v2",
	Short: "Managed OpenStack projects and users",
	Long:  `Managed OpenStack projects and users`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	endpoint := os.Getenv("PSOS_ENDPOINT")
	if len(strings.TrimSpace(endpoint)) == 0 {
		fmt.Println("Please define env PSOS_ENDPOINT")
		os.Exit(1)
	}

	user := os.Getenv("PSOS_USERNAME")
	if len(strings.TrimSpace(user)) == 0 {
		fmt.Println("Please define env PSOS_USERNAME")
		os.Exit(1)
	}

	password := os.Getenv("PSOS_PASSWORD")
	if len(strings.TrimSpace(password)) == 0 {
		fmt.Println("Please define env PSOS_PASSWORD")
		os.Exit(1)
	}

	var err error
	psOsClient, err = psos.Login(endpoint, user, password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
