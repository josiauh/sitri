/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var successColor = color.New(color.FgHiGreen).Add(color.Bold)
var errorColor = color.New(color.FgHiRed).Add(color.Bold)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a Sitri project.",
	Long: `Create a Sitri project.
	
	This will also initialize a git repository, unless specified.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()

		if err != nil {
			errorColor.Println("❌ An error occured when getting the directory you want to initalize!")
			errorColor.Println(err)
			return
		}
		println("🍊 Initializing a Sitri project...")
		println("Initialization at " + cwd)
		os.Mkdir(".sitri", os.ModeDir)
		os.Mkdir(".sitri/commits", os.ModeDir)
		file, ferr := os.Create(".sitri/projectInfo")

		if ferr != nil {
			errorColor.Println("❌ Could not create project info!")
			errorColor.Println(ferr)
			return
		}

		var noGit = contains(args, "-g")
		if !noGit {
			noGit = contains(args, "-noGit")
		}
		if !noGit {
			exec.Command("git", "init", cwd)
		}
		forGit := strconv.FormatBool(noGit)
		file.WriteString(fmt.Sprintf(`[sitriInfo v1]
gitDisabled: %s
		`, forGit))
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
	initCmd.Flags().BoolP("noGit", "g", false, "Don't initialize a git repo.")
}
