package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/JoshVarga/svgparser"
	"github.com/spf13/cobra"
)

// svgCmd represents the svg command
var svgCmd = &cobra.Command{
	Use:   "svg",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: svgCheck,
}

func init() {
	rootCmd.AddCommand(svgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// svgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// svgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func expandGlobs(args []string) ([]string, error) {
	files := []string{}

	for _, arg := range args {
		argfiles, _ := filepath.Glob(arg)
		files = append(files, argfiles...)
	}

	return files, nil
}

func basicTests(filepath string) {
	fi, err := os.Stat(filepath)
	if err != nil {
		fmt.Printf("ERROR: stat failed\n")
		return
	}
	if fi.Size() < minSize {
		fmt.Printf("ERROR: too small: %d < %d\n", fi.Size(), minSize)
		return
	}
	if fi.Size() > maxSize {
		fmt.Printf("ERROR: too big: %d > %d\n", fi.Size(), minSize)
		return
	}
}

func svgCheck(cmd *cobra.Command, args []string) {

	paths, _ := expandGlobs(args)

	for _, filepath := range paths {
		fmt.Printf("INFO: file=%s\n", filepath)
		basicTests(filepath)

		bytes, readErr := ioutil.ReadFile(filepath)
		if readErr != nil {
			fmt.Printf("ERROR: unable to read %s: %s\n", filepath, readErr)
			continue
		}
		text := string(bytes)

		rootElement, parseErr := svgparser.Parse(strings.NewReader(text), false)
		if parseErr != nil {
			fmt.Printf("ERROR: unable to parse %s: %s\n", filepath, parseErr)
			continue
		}
		fmt.Printf("SVG width: %s\n", rootElement.Attributes["width"])

	}
}
