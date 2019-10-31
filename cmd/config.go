package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var config = &cobra.Command{
	Use:   "config",
	Short: "Configure mustang's credentials",
	Run: func(cmd *cobra.Command, args []string) {
		collector := cmd.Flag("collector").Value.String()
		profile := cmd.Flag("profile").Value.String()
		fmt.Println("collector:", collector)
		fmt.Println("profile:", profile)

		f, err := os.OpenFile(profile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Please provide the path to the profile file as well using using ./mustang config --profile /home/rvf/.profile")
			fmt.Println("Error: ", err)
		}
		defer f.Close()
		if _, err := f.WriteString(fmt.Sprintf("export ARGS=%s", collector)); err != nil {
			fmt.Println(err)
		}
		return
	},
}
