/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var bsuccessColor = color.New(color.FgHiGreen).Add(color.Bold)
var berrorColor = color.New(color.FgHiRed).Add(color.Bold)

// ghPublishCmd represents the ghPublish command
var ghPublishCmd = &cobra.Command{
	Use:   "ghPublish",
	Short: "Publish your repository to github. ",
	Long:  `Publish your repository to github. If a non-git project is detected, it will make a git repo.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Publishing to github...")
		sitriInfoB, err := os.ReadFile(".sitri/projectInfo")
		sitriInfoS := string(sitriInfoB)
		sitriInfo := strings.Split(sitriInfoS, "\n")
		if err != nil {
			berrorColor.Printf("❌ Couldn't get the project info! Try re-initializing your Sitri project.")
			os.Exit(1)
		}
		if sitriInfo[1] == "gitDisabled: true" {
			exec.Command("\"C:/Program Files/Git/cmd/git.exe\"", "init")
		}
		exec.Command("gh", "repo", "create", "--public", "--push", "--source", ".")
	},
}

func init() {
	ghCmd.AddCommand(ghPublishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ghPublishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ghPublishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
