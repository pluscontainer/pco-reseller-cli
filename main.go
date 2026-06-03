/*
Copyright © 2022 PlusServer GmbH
*/
package main

import (
	"fmt"
	"os"

	"github.com/pluscontainer/pco-reseller-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
