package cmd

import (
	"bytes"
	"image/jpeg"

	"github.com/spf13/cobra"
)

var (
	jpegHeight Range
	jpegWidth  Range
)

// jpegCmd represents the jpeg command
var jpegCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "jpeg",
	Short: "test JPEG images",
	Long:  `Validate that your JPEG files are valid`,
	RunE:  makeFileCommand(jpegCheck),
}

func init() {
	rootCmd.AddCommand(jpegCmd)

	jpegCmd.Flags().Var(&jpegHeight, "height", "Range of allowed JPEG heights")
	jpegCmd.Flags().Var(&jpegWidth, "width", "Range of allowed JPEG widths")
}

func jpegCheck(f *FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.recordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	image, parseErr := jpeg.Decode(bytes.NewReader(data))

	if parseErr != nil {
		f.recordResult("jpegParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}

	bounds := image.Bounds()

	if jpegWidth.Exists() {
		width := bounds.Max.X - bounds.Min.X
		f.recordResult("jpegWidth", jpegWidth.Check(uint64(width)), map[string]interface{}{
			"desiredWidth": jpegWidth.String(),
			"actualWidth":  width,
		})
	}

	if jpegHeight.Exists() {
		height := bounds.Max.Y - bounds.Min.Y
		f.recordResult("jpegHeight", jpegHeight.Check(uint64(height)), map[string]interface{}{
			"desiredHeight": jpegHeight.String(),
			"actualheight":  height,
		})
	}
}
