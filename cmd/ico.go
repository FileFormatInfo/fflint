package cmd

import (
	"bytes"

	"github.com/zyxar/image2ascii/ico"
	"github.com/spf13/cobra"
)

// icoCmd represents the ico command
var icoCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "ico",
	Short: "test ico files",
	Long:  `Validate that your ico files are valid`,
	RunE:  makeFileCommand(icoCheck),
}

func init() {
	rootCmd.AddCommand(icoCmd)
}

func icoCheck(f *FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.recordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	_, parseErr := ico.DecodeAll(bytes.NewReader(data))

	if parseErr != nil {
		f.recordResult("icoParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}
}
