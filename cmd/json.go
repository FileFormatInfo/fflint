package cmd

import (
	"bytes"

	"github.com/antchfx/jsonquery"
	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "json",
	Short: "test json files",
	Long:  `Validate that your json files are valid`,
	RunE:  makeFileCommand(jsonCheck),
}

func init() {
	rootCmd.AddCommand(jsonCmd)
}

func jsonCheck(f *FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.recordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	_, parseErr := jsonquery.Parse(bytes.NewReader(data))

	if parseErr != nil {
		f.recordResult("jsonParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}
}
