/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var asuccessColor = color.New(color.FgHiGreen).Add(color.Bold)
var aerrorColor = color.New(color.FgHiRed).Add(color.Bold)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears commits stored locally.",
	Long:  `Clears commits stored locally.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🗑️🍊 Trashing your commits...")
		_, err := os.Stat(".sitri")
		if os.IsNotExist(err) {
			aerrorColor.Printf("❌ This isn't a Sitri project! Make sure the directory you are running this in is.")
		}

	},
}

func init() {
	commitCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
