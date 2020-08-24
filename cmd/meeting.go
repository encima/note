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
type Meeting struct {
	With, Purpose   string
	On time.Time
	Attendees []string
	Links []string
}

const meetTmplFile = "./meeting/general.md"

var meetTmpl *template.Template

// bookCmd represents the book command
var meetCmd = &cobra.Command{
	Use:   "meet",
	Short: "Creates a template for creating meeting notes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		meet, err := promptForMeet()
		if err != nil {
			log.Fatalf("error prompting for meeting: %v, aborting", err)
		}
		createMeeting(meet, "$HOME/Dropbox/me/notes/aiven/meetings/")


	},
}

func init() {
	rootCmd.AddCommand(meetCmd)
	box := packr.NewBox("../templates")
	indexTmpl = template.Must(template.New("meet.md").Parse(box.String(meetTmplFile)))
}

func promptForMeet() (*Meeting, error) {
	meet := &Meeting{
		With:     prompt("title: "),
		Purpose:     prompt("status: "),
		On: time.Now(),
		Attendees:   repeatPrompt("Attendees: "),
		Links:   repeatPrompt("Links: "),
	}

	return meet, nil
}

func createMeeting(meeting *Meeting, dir string) error {
	idxf, err := os.Create(os.ExpandEnv(fmt.Sprintf("%s/%s-%s.md", dir, meeting.On, meeting.With)))
	if err != nil {
		return err
	}

	defer idxf.Close()
	w := bufio.NewWriter(idxf)
	defer w.Flush()
	return indexTmpl.Execute(w, meeting)
}


