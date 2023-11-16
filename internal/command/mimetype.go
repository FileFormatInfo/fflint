package command

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

var (
	mimetypeCounterMap   map[string]int //LATER: this should be in command.ctx?
	mimetypeReport       bool
	mimetypeAllowUnknown bool
)

// mimetypeCmd represents the mimetype command
var mimetypeCmd = &cobra.Command{
	Aliases:  []string{"mt", "filetype"},
	Args:     cobra.MinimumNArgs(1),
	Use:      "mimetype [options] files...",
	Short:    "Validate (or report) MIME content types",
	Long:     `Check and report on the MIME content types in use`,
	Example:  `Content type detection uses the Go standard library [http.DetectContentType](https://golang.org/pkg/net/http/#DetectContentType) function.`,
	PreRunE:  mimetypeReportInit,
	RunE:     shared.MakeFileCommand(mimetypeCheck),
	PostRunE: mimetypeReportRun,
}

func AddMimeTypeCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(mimetypeCmd)

	mimetypeCmd.Flags().BoolVar(&mimetypeReport, "report", true, "Print summary report (default is true)")
	mimetypeCmd.Flags().BoolVar(&mimetypeAllowUnknown, "allowUnknown", true, "Allow application/octet-stream")

	//LATER: extension matches mimetype
}

func mimetypeCheck(fc *shared.FileContext) {

	bytes, readErr := fc.ReadFile()
	if readErr != nil {
		fc.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	//LATER: alterate library [h2non/filetype](https://github.com/h2non/filetype)
	//LATER: alterate library [gabriel-vasile/mimetype](https://github.com/gabriel-vasile/mimetype)
	// https://golang.org/pkg/net/http/#DetectContentType
	mimetype := http.DetectContentType(bytes)

	mimetypeCounterMap[mimetype]++

	if !mimetypeAllowUnknown {
		fc.RecordResult("mimetypeAllowUnknown", mimetype != "application/octet-stream", map[string]interface{}{
			"actualMimeType": mimetype,
		})
	}
}

func mimetypeReportInit(cmd *cobra.Command, args []string) error {
	mimetypeCounterMap = make(map[string]int)

	return nil
}

func mimetypeReportRun(cmd *cobra.Command, args []string) error {

	if mimetypeReport {
		if shared.OutputFormat.String() == "json" {
			fmt.Printf("%s\n", shared.EncodeJSON(mimetypeCounterMap))
		} else if shared.OutputFormat.String() == "text" {
			keys := []string{}
			maxKey := 0
			maxValue := 0
			for key, value := range mimetypeCounterMap {
				keys = append(keys, key)
				if len(key) > maxKey {
					maxKey = len(key)
				}
				if value > maxValue {
					maxValue = value
				}
			}

			sort.Strings(keys)
			format := fmt.Sprintf("%%-%ds : %%%dd\n", maxKey, len(fmt.Sprintf("%d", maxValue)))

			for _, key := range keys {
				fmt.Printf(format, key, mimetypeCounterMap[key])
			}
		} else if shared.OutputFormat.String() == "markdown" {
			keys := maps.Keys(mimetypeCounterMap)
			sort.Strings(keys)
			table := tablewriter.NewWriter(os.Stdout)
			table.SetAutoFormatHeaders(false)
			table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT})
			table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
			table.SetCenterSeparator("|")

			table.SetHeader([]string{"Content Type", "Count"})
			total := 0
			for _, key := range keys {
				table.Append([]string{key, strconv.Itoa(mimetypeCounterMap[key])})
				total += mimetypeCounterMap[key]
			}
			table.Append([]string{"Total:", strconv.Itoa(total)})
			table.Render()
		}

	}

	return nil
}
