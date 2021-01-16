package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var (
	caseSensitive bool
	counterMap    map[string]int //LATER: this should be in command.ctx?
	extReport     bool
	extLength     Range
	extAllowEmpty bool
)

// extCmd represents the extension command
var extCmd = &cobra.Command{
	Aliases:  []string{"extension", "extensions"},
	Args:     cobra.MinimumNArgs(1),
	Use:      "ext [flags] filespec [filespec...]",
	Short:    "test/report file extensions",
	Long:     ``,
	PreRunE:  extensionReportInit,
	RunE:     makeFileCommand(extCheck),
	PostRunE: extensionReportRun,
}

func init() {
	rootCmd.AddCommand(extCmd)

	extCmd.Flags().BoolVar(&caseSensitive, "caseSensitive", false, "Case sensitive (default is false)")
	extCmd.Flags().BoolVar(&extReport, "report", true, "Print summary report (default is true)")
	extCmd.Flags().Var(&extLength, "extlen", "Range of allowed extension lengths")
	extCmd.Flags().BoolVar(&extAllowEmpty, "allowEmpty", true, "Allow files without an extension")
	//LATER: allowed: list of acceptable extension
	//LATER: length: range
	//LATER: allowNone
	//LATER: minCount/maxCount: per-extension min/max
}

// differs from standard golang about .dotfiles and return value doesn't include dot
func getExt(path string) string {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			if i == 0 || os.IsPathSeparator(path[i-1]) {
				return ""
			}
			return path[i+1:]
		}
	}
	return ""
}

func extCheck(fc *FileContext) {

	ext := getExt(fc.FilePath)
	if !caseSensitive {
		ext = strings.ToLower(ext)
	}
	counterMap[ext]++

	if ext == "" {
		fc.recordResult("extAllowEmpty", extAllowEmpty, nil)
	} else if extLength.Exists() {
		fc.recordResult("extLength", extLength.Check(uint64(len(ext))), map[string]interface{}{
			"desiredLength": extLength.String(),
			"actualLength":  len(ext),
		})
	}
}

func extensionReportInit(cmd *cobra.Command, args []string) error {
	counterMap = make(map[string]int)

	return nil
}

func extensionReportRun(cmd *cobra.Command, args []string) error {

	if extReport {
		if outputFormat == "json" {
			fmt.Printf("%s\n", encodeJSON(counterMap))
		} else {
			keys := []string{}
			maxKey := 0
			maxValue := 0
			for key, value := range counterMap {
				keys = append(keys, key)
				if len(key) > maxKey {
					maxKey = len(key)
				}
				if value > maxValue {
					maxValue = value
				}
			}

			sort.Strings(keys)
			format := fmt.Sprintf("%%-%ds: %%%dd\n", maxKey, len(fmt.Sprintf("%d", maxValue)))

			for _, key := range keys {
				fmt.Printf(format, key, counterMap[key])
			}
		}
	}

	return nil
}
