package cmd

import (
	"bytes"

	"github.com/antchfx/xmlquery"
	"github.com/spf13/cobra"
)

// xmlCmd represents the xml command
var xmlCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "xml",
	Short: "test xml files",
	Long:  `Validate that your xml files are valid`,
	RunE:  makeFileCommand(xmlCheck),
}

func init() {
	rootCmd.AddCommand(xmlCmd)
}

func xmlCheck(f *FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.recordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	_, parseErr := xmlquery.Parse(bytes.NewReader(data))

	if parseErr != nil {
		f.recordResult("xmlParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}
}
