// Copyright © 2019 Sean K Smith <ssmith2347@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// TODO - Increase test coverage
// TODO - Update source control
// TODO - Add config file
//         - home directory, etc.
// TODO - Add github integration

package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

var cfgFile string
var s *bufio.Scanner

const filePermissions = 0755

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "note",
	Short: "Creates a basic markdown template for notes",
	Long: `
	Creates a series of linked markdown files in a simple
	template useful for detailed note taking. Store the project
	in a git repository for easy access and tracking fo your
	notes.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Settings represents application settings supplied by the user
type Settings struct {
	Home    string
	BlogDir string `json:"blogdir"`
}

var settings *Settings

func loadSettings() {
	usr, err := user.Current()

	if err != nil {
		log.Fatal("Unable to acquire user")
	}

	settings = &Settings{
		Home: usr.HomeDir,
	}
}
