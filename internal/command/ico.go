package command

import (
	"bytes"

	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
	"github.com/zyxar/image2ascii/ico"
)

// icoCmd represents the ico command
var icoCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "ico [options] files...",
	Short: "Validate icons",
	Long:  `Validate that your ico files are valid`,
	RunE:  shared.MakeFileCommand(icoCheck),
}

func AddIcoCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(icoCmd)

	//LATER: sizes flag (array of ints)
}

func icoCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	_, parseErr := ico.DecodeAll(bytes.NewReader(data))

	if parseErr != nil {
		f.RecordResult("icoParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}
}
