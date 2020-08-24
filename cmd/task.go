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

package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
)


// Book represents the book getting created
type Task struct {
	Name, Desc, Status, For, From, Area   string
	Created, Due time.Time
	Links []string
}

const taskTmplFile = "./task/task.md"

var taskTmpl *template.Template

// bookCmd represents the book command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Creates a template for creating tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		task, err := promptForTask()
		if err != nil {
			log.Fatalf("error prompting for task: %v, aborting", err)
		}
		createTask(task, fmt.Sprintf("$HOME/Dropbox/me/tasks/%s/%s/", task.Area, task.Status))


	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
	box := packr.NewBox("../templates")
	indexTmpl = template.Must(template.New("task.md").Parse(box.String(taskTmplFile)))
}

func promptForTask() (*Task, error) {
	task := &Task{
		Name:     prompt("title: "),
		Status:     prompt("status: "),
		For:  prompt("for: "),
		From: prompt("from: "),
		Created: time.Now(),
		Area: prompt("area: "),
		Links:   repeatPrompt("links: "),
	}

	return task, nil
}

func createTask(task *Task, dir string) error {
	idxf, err := os.Create(os.ExpandEnv(fmt.Sprintf("%s/%s.md", dir, task.Name)))
	if err != nil {
		return err
	}

	defer idxf.Close()
	w := bufio.NewWriter(idxf)
	defer w.Flush()
	return indexTmpl.Execute(w, task)
}


