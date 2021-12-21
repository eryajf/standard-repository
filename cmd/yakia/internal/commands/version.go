package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "v0.1.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Yakia version",
	Long:  "Show Yakia version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Yakia version " + cmd.Root().Version)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
