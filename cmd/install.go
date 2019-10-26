package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var install = &cobra.Command{
	Use:   "install",
	Short: "install node_exporter on the machine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing node_Exporter on : ", runtime.GOOS)
	},
}
