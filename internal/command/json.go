package command

import (
	"encoding/json"

	"github.com/FileFormatInfo/fflint/internal/shared"
	//"github.com/antchfx/jsonquery"
	"github.com/spf13/cobra"
)

var (
	jsonSchemaValidator shared.SchemaOptions
	//jsonSchemaLocation  string
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Args:    cobra.MinimumNArgs(1),
	Use:     "json [options] files...",
	Short:   "Validate JSON files",
	Long:    `Check that your JSON files are valid`,
	PreRunE: jsonInit,
	RunE:    shared.MakeFileCommand(jsonCheck),
}

func AddJsonCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(jsonCmd)

	jsonSchemaValidator.AddFlags(jsonCmd)
	//LATER: whitespace: canonical/none/any
}

func jsonCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	var jsonData any
	parseErr := json.Unmarshal(data, &jsonData)

	f.RecordResult("jsonParse", parseErr == nil, map[string]interface{}{
		"error": parseErr,
	})
	if parseErr != nil {
		return
	}

	jsonSchemaValidator.Validate(f, jsonData)
}

func jsonInit(cmd *cobra.Command, args []string) error {

	prepErr := jsonSchemaValidator.Prepare()
	if prepErr != nil {
		return prepErr
	}

	return nil
}
