package cmd

import (
	"bytes"
	"fmt"
	"image/png"

	"github.com/spf13/cobra"
)

var (
	pngHeight Range
	pngWidth  Range
)

// pngCmd represents the png command
var pngCmd = &cobra.Command{
	Use:   "png",
	Short: "test png images",
	Long:  `Validate that your png files are valid`,
	Run:   pngCheck,
}

func init() {
	rootCmd.AddCommand(pngCmd)

	pngCmd.Flags().Var(&pngHeight, "height", "Range of allowed PNG heights")
	pngCmd.Flags().Var(&pngWidth, "width", "Range of allowed PNG widths")
}

func pngCheck(cmd *cobra.Command, args []string) {

	files, _ := expandGlobs(args)

	for _, f := range files {
		basicTests(f)

		data, readErr := f.ReadFile()
		if readErr != nil {
			fmt.Printf("ERROR: unable to read %s: %s\n", f.FilePath, readErr)
			continue
		}

		image, parseErr := png.Decode(bytes.NewReader(data))

		if parseErr != nil {
			f.recordResult("pngParse", false, map[string]interface{}{
				"error": parseErr,
			})
			continue
		}

		bounds := image.Bounds()

		if pngWidth.Exists() {
			width := bounds.Max.X - bounds.Min.X
			f.recordResult("pngWidth", pngWidth.Check(uint64(width)), map[string]interface{}{
				"desiredWidth": pngWidth.String(),
				"actualWidth":  width,
			})
		}

		if pngHeight.Exists() {
			height := bounds.Max.Y - bounds.Min.Y
			f.recordResult("pngHeight", pngHeight.Check(uint64(height)), map[string]interface{}{
				"desiredHeight": pngHeight.String(),
				"actualheight":  height,
			})
		}
	}
}
