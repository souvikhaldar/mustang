package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "mustang",
	Short: "mustang set's up your machine for sending metrics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mustang is here!")
	},
}

func init() {
	RootCmd.AddCommand(install)
	RootCmd.AddCommand(config)
	config.PersistentFlags().String("service", "", "full name of the systemd service to monitor. Eg.-  http-server.service")
	config.PersistentFlags().String("profile", "~/.profile", "path to the profile file. Eg.- ~/.profile, /root/svkhldr/.profile")
}
