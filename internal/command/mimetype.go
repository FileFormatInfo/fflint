package command

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
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
	Use:      "mimetype [flags] filespec [filespec...]",
	Short:    "test/report mime types",
	Long:     ``,
	PreRunE:  mimetypeReportInit,
	RunE:     shared.MakeFileCommand(mimetypeCheck),
	PostRunE: mimetypeReportRun,
}

func AddMimeTypeCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(mimetypeCmd)

	mimetypeCmd.Flags().BoolVar(&mimetypeReport, "report", true, "Print summary report (default is true)")
	mimetypeCmd.Flags().BoolVar(&mimetypeAllowUnknown, "allowUnknown", true, "Allow application/octet-stream")
}

func mimetypeCheck(fc *shared.FileContext) {

	bytes, readErr := fc.ReadFile()
	if readErr != nil {
		fc.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}
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
		if shared.OutputFormat == "json" {
			fmt.Printf("%s\n", shared.EncodeJSON(mimetypeCounterMap))
		} else {
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
		}
	}

	return nil
}
