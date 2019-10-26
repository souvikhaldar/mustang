package main

import (
	"os"

	"github.com/souvikhaldar/mustang/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
