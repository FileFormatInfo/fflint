package command

import (
	"bytes"

	"github.com/antchfx/jsonquery"
	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "json",
	Short: "Validate JSON files",
	Long:  `Check that your JSON files are valid`,
	RunE:  shared.MakeFileCommand(jsonCheck),
}

func AddJsonCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(jsonCmd)

	//LATER: whitespace: canonical/none/any
	//LATER: schema (https://github.com/xeipuuv/gojsonschema)
}

func jsonCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	_, parseErr := jsonquery.Parse(bytes.NewReader(data))

	if parseErr != nil {
		f.RecordResult("jsonParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}
}
