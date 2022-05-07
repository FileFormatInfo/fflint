package command

import (
	"bytes"
	"image/png"

	"github.com/fileformat/badger/internal/argtype"
	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
)

var (
	pngHeight argtype.Range
	pngWidth  argtype.Range
)

// pngCmd represents the png command
var pngCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "png",
	Short: "test png images",
	Long:  `Validate that your png files are valid`,
	RunE:  shared.MakeFileCommand(pngCheck),
}

func AddPngCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(pngCmd)

	pngCmd.Flags().Var(&pngHeight, "height", "Range of allowed PNG heights")
	pngCmd.Flags().Var(&pngWidth, "width", "Range of allowed PNG widths")
}

func pngCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	image, parseErr := png.Decode(bytes.NewReader(data))

	if parseErr != nil {
		f.RecordResult("pngParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}

	bounds := image.Bounds()

	if pngWidth.Exists() {
		width := bounds.Max.X - bounds.Min.X
		f.RecordResult("pngWidth", pngWidth.Check(uint64(width)), map[string]interface{}{
			"desiredWidth": pngWidth.String(),
			"actualWidth":  width,
		})
	}

	if pngHeight.Exists() {
		height := bounds.Max.Y - bounds.Min.Y
		f.RecordResult("pngHeight", pngHeight.Check(uint64(height)), map[string]interface{}{
			"desiredHeight": pngHeight.String(),
			"actualheight":  height,
		})
	}
}
