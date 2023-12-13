/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// ghCmd represents the gh command
var ghCmd = &cobra.Command{
	Use:   "git",
	Short: "A series of Git commands.",
	Long:  `A series of Git commands.`,
}

func init() {
	rootCmd.AddCommand(ghCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ghCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ghCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
