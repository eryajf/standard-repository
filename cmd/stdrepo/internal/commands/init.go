/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize directories and files",
	Long: `Initialize directories and files`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		// fmt.Println(cmd.Flags().GetString("projecName"))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
	// initCmd.Flags().StringP("projecName", "n","", "Set the project name.")
}
