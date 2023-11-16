package command

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/antchfx/jsonquery"
	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
)

var (
	jsonSchemaLocation string
	jsonSchema         *gojsonschema.Schema
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

	jsonCmd.Flags().StringVar(&jsonSchemaLocation, "schema", "", "JSON Schema to validate against") //LATER: link to docs about embedded ones

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

	if jsonSchema != nil {
		result, validateErr := jsonSchema.Validate(gojsonschema.NewStringLoader(string(data)))
		if validateErr != nil {
			f.RecordResult("jsonSchemaRun", false, map[string]interface{}{
				"error": validateErr.Error(),
			})
		} else {
			f.RecordResult("jsonSchemaValidate", result.Valid(), map[string]interface{}{
				"errors": result.Errors(),
			})
		}
	}

}

func jsonInit(cmd *cobra.Command, args []string) error {

	if jsonSchemaLocation == "" {
		return nil
	}

	// work with local file urls
	jsonUrl, urlParseErr := url.Parse(jsonSchemaLocation)
	if urlParseErr != nil {
		return urlParseErr
	}

	// allow relative local file schemas
	if jsonUrl.Scheme == "" {
		jsonUrl.Scheme = "file"
		jsonPath, pathErr := filepath.Abs(jsonUrl.Path)
		if pathErr != nil {
			return pathErr
		}
		jsonUrl.Path = jsonPath
		newLocation := jsonUrl.String()
		if shared.Debug {
			fmt.Fprintf(os.Stderr, "DEBUG: canonicalizing schema path from '%s' to '%s'\n", jsonSchemaLocation, newLocation)
		}
		jsonSchemaLocation = newLocation
	}

	jsonSchemaLoader := gojsonschema.NewReferenceLoader(jsonSchemaLocation)
	var schemaErr error
	jsonSchema, schemaErr = gojsonschema.NewSchema(jsonSchemaLoader)
	return schemaErr
}
