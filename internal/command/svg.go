package command

import (
	"strconv"
	"strings"

	"github.com/JoshVarga/svgparser"
	"github.com/fileformat/badger/internal/argtype"
	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
)

var (
	svgHeight  argtype.Range
	svgWidth   argtype.Range
	svgViewBox argtype.DecimalRangeArray
)

// svgCmd represents the svg command
var svgCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "svg",
	Short: "Validate SVG images",
	Long:  `Check that SVG files are error free and (optionally) don't have any undesirable things in them.`,
	RunE:  shared.MakeFileCommand(svgCheck),
}

func AddSvgCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(svgCmd)

	svgCmd.Flags().Var(&svgHeight, "height", "Range of allowed SVG heights")
	svgCmd.Flags().Var(&svgViewBox, "viewbox", "Ranges of allowed SVG viewBox values")
	svgCmd.Flags().Var(&svgWidth, "width", "Range of allowed SVG widths")
	//LATER: no text
	//LATER: raster inclusions: none/embedded/linked/any
	//LATER: no external links
	//LATER: current color
	//LATER: xml namespace
	//LATER: no additional namespaces
	//LATER: no foreign objects
	//LATER: dimension units on width/height/viewBox: true/false/listOfAcceptable
	//LATER: font
	//LATER: meta
	//LATER: optimized (and/or pretty?)

}

func svgCheck(f *shared.FileContext) {

	bytes, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}
	text := string(bytes)

	rootElement, parseErr := svgparser.Parse(strings.NewReader(text), false)
	if parseErr != nil {
		f.RecordResult("svgParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}

	if svgWidth.Exists() {
		widthStr := rootElement.Attributes["width"]
		width, err := strconv.ParseUint(widthStr, 10, 64)
		if err != nil {
			f.RecordResult("svgWidth", false, map[string]interface{}{
				"error": err,
				"width": widthStr,
			})
		} else {
			f.RecordResult("svgWidth", svgWidth.Check(width), map[string]interface{}{
				"desiredWidth": svgWidth.String(),
				"actualWidth":  width,
			})
		}
	}

	if svgHeight.Exists() {
		heightStr := rootElement.Attributes["height"]
		height, err := strconv.ParseUint(heightStr, 10, 64)
		if err != nil {
			f.RecordResult("svgHeight", false, map[string]interface{}{
				"error":  err,
				"height": heightStr,
			})
		} else {
			f.RecordResult("svgHeight", svgHeight.Check(height), map[string]interface{}{
				"desiredHeight": svgHeight.String(),
				"actualheight":  height,
			})
		}
	}

	if svgViewBox.Exists() {
		viewBoxStr := rootElement.Attributes["viewBox"]
		f.RecordResult("svgViewBox", svgViewBox.CheckString(viewBoxStr, " "), map[string]interface{}{
			"actualViewBox":   viewBoxStr,
			"expectedViewBox": svgViewBox.String(),
		})
	}

}
