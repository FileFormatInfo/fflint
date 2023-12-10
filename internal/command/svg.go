package command

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/FileFormatInfo/fflint/internal/argtype"
	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/JoshVarga/svgparser"
	"github.com/spf13/cobra"
)

var (
	svgHeight       argtype.Range
	svgWidth        argtype.Range
	svgViewBox      argtype.DecimalRangeArray
	svgNamespace    bool
	svgNamespaces   []string
	svgNamespaceSet map[string]bool
	svgText         bool
	svgForeign      bool
	svgImage        = argtype.NewStringSet("Images", "none", []string{"any", "embedded", "external", "none"})
)

func AddSvgCommand(rootCmd *cobra.Command) {
	var svgCmd = &cobra.Command{
		Args:    cobra.MinimumNArgs(1),
		Use:     "svg [options] files...",
		Short:   "Validate SVG images",
		Long:    `Check that SVG files are error free and (optionally) don't have any undesirable things in them.`,
		PreRunE: svgPrepare,
		RunE:    shared.MakeFileCommand(svgCheck),
	}

	svgCmd.Flags().Var(&svgHeight, "height", "Range of allowed SVG heights")
	svgCmd.Flags().Var(&svgViewBox, "viewbox", "Ranges of allowed SVG viewBox values")
	svgCmd.Flags().Var(&svgWidth, "width", "Range of allowed SVG widths")
	svgCmd.Flags().BoolVar(&svgNamespace, "namespace", true, "Check namespaces")
	svgCmd.Flags().StringSliceVar(&svgNamespaces, "namespaces", []string{}, "Additional namespaces allowed when checking namespaces (`*` for all)")
	svgCmd.Flags().BoolVar(&svgText, "text", false, "Allow text nodes")
	svgCmd.Flags().BoolVar(&svgForeign, "foreign", false, "Allow foreign objects")
	svgCmd.Flags().Var(&svgImage, "images", "Embedded image control "+svgImage.HelpText())
	//LATER: raster inclusions: none/embedded/linked/any (https://github.com/svg/svgo/blob/main/plugins/removeRasterImages.js)
	//LATER: off-canvas paths (https://github.com/svg/svgo/blob/main/plugins/removeOffCanvasPaths.js)
	//LATER: external links: OptionalBool
	//LATER: gzip'ed: OptionalBool
	//LATER: current color
	//LATER: xml namespace
	//LATER: no additional namespaces
	//LATER: no foreign objects
	//LATER: dimension units on width/height/viewBox: true/false/listOfAcceptable
	//LATER: font
	//LATER: meta
	//LATER: optimized (and/or pretty?)

	rootCmd.AddCommand(svgCmd)
}

func svgCheck(f *shared.FileContext) {

	raw, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}
	text := string(raw)

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

	if svgNamespace {

		namespaces, _ := shared.GetNamespaces(bytes.NewReader(raw))
		//fmt.Printf("namespace: %v", namespaces)
		f.RecordResult("svgNamespace", namespaces.Default == "http://www.w3.org/2000/svg", map[string]interface{}{
			"namespace": namespaces.Default,
		})

		if len(svgNamespaces) == 0 {
			f.RecordResult("svgNoAdditionalNamespaces", len(namespaces.Additional) == 0, map[string]interface{}{
				"namespaces": namespaces.Additional,
			})
		} else if len(svgNamespaces) == 1 && svgNamespaces[0] == "*" {
			// no check
		} else {
			for key, value := range namespaces.Additional {
				_, keyExists := svgNamespaceSet[key]
				_, valueExists := svgNamespaceSet[value]
				f.RecordResult("svgAdditionalNamespaces", keyExists || valueExists, map[string]interface{}{
					"namespaceUrl":   value,
					"namespaceValue": key,
				})
			}
		}
	}

	if !svgText {
		textNodes := rootElement.FindAll("text")
		if len(textNodes) > 0 {
			f.RecordResult("svgText", false, map[string]interface{}{
				"textNodeCount": len(textNodes),
				"textContent":   getLimitedContent(textNodes, 80),
			})
		}
	}

	/* doesn't seem to work??? */
	if !svgForeign {
		foNodes := rootElement.FindAll("foreignObject")
		if len(foNodes) > 0 {
			f.RecordResult("svgForeignObject", false, map[string]interface{}{
				"foreignObjectNodeCount": len(foNodes),
			})
		}
	}

	if svgImage.String() != "any" {
		imageNodes := rootElement.FindAll("image")
		if len(imageNodes) > 0 {
			if svgImage.String() == "none" {
				f.RecordResult("svgImage", false, map[string]interface{}{
					"imageNodeCount": len(imageNodes),
				})
			}
		}

	}
}

// LATER: limit total length of content add ellipsis
func getContent(contentCache []string, elist []*svgparser.Element) []string {
	if contentCache == nil {
		contentCache = []string{}
	}

	if len(elist) == 0 {
		return contentCache
	}

	for _, el := range elist {
		contentCache = append(contentCache, el.Content)
		contentCache = getContent(contentCache, el.Children)
	}

	return contentCache
}
func getLimitedContent(elist []*svgparser.Element, limit int) string {

	content := getContent(nil, elist)

	retVal := strings.Join(content, " ") //LATER: custom loop with early break
	if len(retVal) > limit {
		retVal = retVal[:limit-1] + "\u2026"
	}
	return retVal
}

func svgPrepare(cmd *cobra.Command, args []string) error {
	if svgNamespace {
		svgNamespaceSet = make(map[string]bool)
		for _, key := range svgNamespaces {
			svgNamespaceSet[key] = true
		}
	}
	return nil
}
