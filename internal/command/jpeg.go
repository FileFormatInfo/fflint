package command

import (
	"bytes"
	"image/jpeg"

	"github.com/fileformat/badger/internal/argtype"
	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
)

var (
	jpegHeight argtype.Range
	jpegWidth  argtype.Range
)

// jpegCmd represents the jpeg command
var jpegCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "jpeg",
	Short: "Validate JPEG images",
	Long:  `Validate that your JPEG files are valid`,
	RunE:  shared.MakeFileCommand(jpegCheck),
}

func AddJpegCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(jpegCmd)

	jpegCmd.Flags().Var(&jpegHeight, "height", "Range of allowed JPEG heights")
	jpegCmd.Flags().Var(&jpegWidth, "width", "Range of allowed JPEG widths")
	//LATER: aspect ratio (range)
	//LATER: metadata: https://github.com/dsoprea/go-exif or https://github.com/rwcarlsen/goexif
	//LATER: color profile
}

func jpegCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	image, parseErr := jpeg.Decode(bytes.NewReader(data))

	if parseErr != nil {
		f.RecordResult("jpegParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}

	bounds := image.Bounds()

	if jpegWidth.Exists() {
		width := bounds.Max.X - bounds.Min.X
		f.RecordResult("jpegWidth", jpegWidth.Check(uint64(width)), map[string]interface{}{
			"desiredWidth": jpegWidth.String(),
			"actualWidth":  width,
		})
	}

	if jpegHeight.Exists() {
		height := bounds.Max.Y - bounds.Min.Y
		f.RecordResult("jpegHeight", jpegHeight.Check(uint64(height)), map[string]interface{}{
			"desiredHeight": jpegHeight.String(),
			"actualheight":  height,
		})
	}
}
