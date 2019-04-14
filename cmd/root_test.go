package cmd

import (
	"bufio"
	"os"
	"testing"
)

func TestFormatDirName(t *testing.T) {
	const exp = "some--directory--namedir-name"
	const inp = `  Some  Directory -Name

	dir-name`

	d, err := formatDirName(inp)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if d != exp {
		t.Errorf(`formatDirName("%s") returned "%s", want %s`, inp, d, exp)
	}
}

func TestFormatDirNameInvalid(t *testing.T) {
	_, err := formatDirName("")
	if err == nil {
		t.Errorf("expected error not returned")
	}
}

func TestRootCommand(t *testing.T) {
	// rootCmd.Flags().AddFlag(&pflag.Flag{Name: "title", Value: "A Long St"})

	if err := rootCmd.Execute(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCreateDirectories(t *testing.T) {
	const testDir = "test-dir"

	if err := createDirs([]string{testDir}); err != nil {
		t.Fatalf("unexpected error creating directories: %v", err)
	}

	if _, err := os.Stat("./" + testDir); os.IsNotExist(err) {
		t.Fatalf("directory not created: %v", err)
	}

	const filepath = "./" + testDir + "/test.txt"
	f, err := os.Create(filepath)
	defer os.RemoveAll("./" + testDir)
	defer f.Close()

	if err != nil {
		t.Fatalf("unable to create file in directory: %v", err)
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString("some data")

	if err != nil {
		t.Errorf("unable to write to file: %v", err)
	}

	w.Flush()

}
