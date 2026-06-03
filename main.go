/*
Copyright © 2022 PlusServer GmbH
*/
package main

import (
	"os"

	"github.com/pluscontainer/pco-reseller-cli/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Errorf("%s", err)
		os.Exit(1)
	}
}
