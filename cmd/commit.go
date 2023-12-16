/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"archive/zip"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sitriproject/sitri/sitrilib"
	"strings"
	"unsafe"

	"github.com/spf13/cobra"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyz1234567890")

func generate(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alphabet[b[i]%byte(len(alphabet))]
	}
	return *(*string)(unsafe.Pointer(&b))
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

var git bool
var message string

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a commit.",
	Long:  `Creates a commit.`,
	Run: func(cmd *cobra.Command, args []string) {
		info, err := sitrilib.GetSitriInfo()
		if err != nil {
			fmt.Printf("Could not get sitri info.")
			os.Exit(1)
		}
		fmt.Println("📦 Commiting your Sitri project...")
		cmtName := generate(6)
		if !git {
			msgFile, err := os.Create(".sitri/commit-msg")
			if err != nil {
				errorColor.Printf("❌ Could not create a commit message file!")
				errorColor.Printf("Full error:")
				errorColor.Printf(err.Error())
				os.Exit(1)
			}
			msgFile.WriteString(message)
			file, err := os.Create(".sitri/commits/commit-" + cmtName + ".zip")
			strignoref, ferr := os.ReadFile(".sitriignore")
			if ferr != nil {
				if os.IsNotExist(ferr) {
					errorColor.Printf("❌ Could not read sitriignore!")
					errorColor.Printf("Full error:")
					errorColor.Printf(ferr.Error())
					os.Exit(1)
				}
			} else {
				strignoref = []byte(``)
			}
			if err != nil {
				errorColor.Printf("❌ Could not create a commit zip!")
				errorColor.Printf("Full error:")
				errorColor.Printf(err.Error())
				os.Exit(1)
			}
			strignore := string(strignoref)
			strignorefiles := strings.Split(strignore, "\n")
			defer file.Close()

			w := zip.NewWriter(file)
			defer w.Close()

			walker := func(path string, info os.FileInfo, err error) error {
				fmt.Printf("Zipping file %#v\n", path)
				fp := strings.Split(path, string(os.PathSeparator))
				name := fp[len(fp)-1]
				if stringInSlice(name, strignorefiles) {
					return nil
				}

				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				// Ensure that `path` is not absolute; it should not start with "/".
				// This snippet happens to work because I don't use
				// absolute paths, but ensure your real-world code
				// transforms path into a zip-root relative path.
				f, err := w.Create(path)
				if err != nil {
					return err
				}

				_, err = io.Copy(f, file)
				if err != nil {
					return err
				}

				return nil
			}
			err = filepath.Walk(".", walker)
			if err != nil {
				fmt.Printf("Could not zip your commit!")
				os.Exit(1)
			}
		}
		if info[1] != "gitDisabled: true" {
			exec.Command("\"C:/Program Files/Git/cmd/git.exe\"", "add", ".")
			exec.Command("\"C:/Program Files/Git/cmd/git.exe\"", "commit", "-m", "\"🍊 Commited from Sitri ("+cmtName+" "+message+")\"")
		}
		fmt.Printf("🍊 Commit created!")
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	commitCmd.Flags().BoolVarP(&git, "git", "g", false, "Only use Git.")
	commitCmd.Flags().StringVarP(&message, "message", "m", "No commit message.", "")
}
