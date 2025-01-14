/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stdrepo",
	Short: "Standard Repository CLI",
	Long:  `Standard Repository CLI`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Standard Repository CLI")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.demo-cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// rootCmd.PersistentFlags().StringP("language","l", "en", "Set the default email.")

	rootCmd.CompletionOptions.HiddenDefaultCmd = true

}
