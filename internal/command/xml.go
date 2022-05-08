package command

import (
	"bytes"

	"github.com/antchfx/xmlquery"
	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
)

// xmlCmd represents the xml command
var xmlCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "xml",
	Short: "Validate XML files",
	Long:  `Checks that your XML files are valid`,
	RunE:  shared.MakeFileCommand(xmlCheck),
}

func AddXmlCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(xmlCmd)
}

func xmlCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	_, parseErr := xmlquery.Parse(bytes.NewReader(data))

	if parseErr != nil {
		f.RecordResult("xmlParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}
}
