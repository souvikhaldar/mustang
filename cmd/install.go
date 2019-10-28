package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var install = &cobra.Command{
	Use:   "install",
	Short: "install node_exporter on the machine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing node_Exporter on : ", runtime.GOOS)
		res,err := exec.Command("/bin/sh", "-c", "sudo apt-get install prometheus-node-exporter").Output()
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(res)
		// if err := c.Run(); err != nil {
		// 	fmt.Println("Error in installing mustang: ", err)
		// }
	},
}
