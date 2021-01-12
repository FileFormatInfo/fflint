package cmd

import (
	"strconv"
	"strings"

	"github.com/JoshVarga/svgparser"
	"github.com/spf13/cobra"
)

var (
	svgHeight Range
	svgWidth  Range
)

// svgCmd represents the svg command
var svgCmd = &cobra.Command{
	Use:   "svg",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: makeFileCommand(svgCheck),
}

func init() {
	rootCmd.AddCommand(svgCmd)

	svgCmd.Flags().Var(&svgHeight, "height", "Range of allowed SVG heights")
	svgCmd.Flags().Var(&svgWidth, "width", "Range of allowed SVG widths")
	//svgCmd.Flags().Var(&svgViewBox, "viewBox", "Ranges of allowed SVG viewBox values")
}

func svgCheck(f FileContext) {

	bytes, readErr := f.ReadFile()
	if readErr != nil {
		f.recordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}
	text := string(bytes)

	rootElement, parseErr := svgparser.Parse(strings.NewReader(text), false)
	if parseErr != nil {
		f.recordResult("svgParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}

	if svgWidth.Exists() {
		widthStr := rootElement.Attributes["width"]
		width, err := strconv.ParseUint(widthStr, 10, 64)
		if err != nil {
			f.recordResult("svgWidth", false, map[string]interface{}{
				"error": err,
				"width": widthStr,
			})
		} else {
			f.recordResult("svgWidth", svgWidth.Check(width), map[string]interface{}{
				"desiredWidth": svgWidth.String(),
				"actualWidth":  width,
			})
		}
	}

	if svgHeight.Exists() {
		heightStr := rootElement.Attributes["height"]
		height, err := strconv.ParseUint(heightStr, 10, 64)
		if err != nil {
			f.recordResult("svgHeight", false, map[string]interface{}{
				"error":  err,
				"height": heightStr,
			})
		} else {
			f.recordResult("svgHeight", svgHeight.Check(height), map[string]interface{}{
				"desiredHeight": svgHeight.String(),
				"actualheight":  height,
			})
		}
	}

	/*
		if svgViewBox.Exists() {
			viewBoxStr := rootElement.Attributes["viewBox"]
			f.recordResult("svgViewBox", true, map[string]interface{}{
				"actualViewBox": viewBoxStr,
			})
		}
	*/
}
