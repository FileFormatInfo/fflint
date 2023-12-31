package command

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/FileFormatInfo/fflint/internal/argtype"
	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

var (
	caseSensitive bool
	counterMap    map[string]int //LATER: this should be in command.ctx?
	extReport     bool
	extLength     argtype.Range
	extAllowEmpty bool
	extAllowed    []string
)

// extCmd represents the extension command
var extCmd = &cobra.Command{
	Aliases:  []string{"extension", "extensions"},
	Args:     cobra.MinimumNArgs(1),
	Use:      "ext [options] files...",
	Short:    "Validate (or report) file extensions",
	Long:     `Check and report on the file extensions in use`,
	PreRunE:  extensionReportInit,
	RunE:     shared.MakeFileCommand(extCheck),
	PostRunE: extensionReportRun,
}

func AddExtCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(extCmd)

	extCmd.Flags().BoolVar(&caseSensitive, "caseSensitive", false, "Case sensitive")
	extCmd.Flags().BoolVar(&extReport, "report", true, "Print summary report")
	extCmd.Flags().Var(&extLength, "length", "Range of allowed extension lengths")
	extCmd.Flags().BoolVar(&extAllowEmpty, "allowEmpty", true, "Allow files without an extension")
	extCmd.Flags().StringSliceVar(&extAllowed, "allowed", []string{}, "Allowed extensions") //LATER: maybe switch to regex?
	//LATER: forbidden: list of unacceptable extensions
	//LATER: allowNone: allow files without an extension
	//LATER: skipDotFiles: skip dot files
	//LATER: skipDotDirs: skip dot directories (i.e. .git)
	//LATER: print the first few files with each extension
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

func extCheck(fc *shared.FileContext) {

	ext := getExt(fc.FilePath)
	if !caseSensitive {
		ext = strings.ToLower(ext)
	}
	counterMap[ext]++

	if ext == "" {
		fc.RecordResult("extAllowEmpty", extAllowEmpty, nil)
	} else if extLength.Exists() {
		fc.RecordResult("extLength", extLength.Check(uint64(len(ext))), map[string]interface{}{
			"desiredLength": extLength.String(),
			"actualLength":  len(ext),
		})
	}

	if len(extAllowed) > 0 {
		fc.RecordResult("extAllowed", contains(extAllowed, ext), map[string]interface{}{
			"ext": ext,
		})
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func extensionReportInit(cmd *cobra.Command, args []string) error {
	counterMap = make(map[string]int)

	return nil
}

func extensionReportRun(cmd *cobra.Command, args []string) error {

	if extReport {
		if shared.OutputFormat.String() == "json" {
			fmt.Printf("%s\n", shared.EncodeJSON(counterMap))
		} else if shared.OutputFormat.String() == "text" {
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
		} else if shared.OutputFormat.String() == "markdown" {
			keys := maps.Keys(counterMap)
			sort.Strings(keys)
			table := tablewriter.NewWriter(os.Stdout)
			table.SetAutoFormatHeaders(false)
			table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT})
			table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
			table.SetCenterSeparator("|")

			table.SetHeader([]string{"Extension", "Count"})
			total := 0
			for _, key := range keys {
				var niceKey string
				if key == "" {
					niceKey = "(empty)"
				} else {
					niceKey = fmt.Sprintf(".%s", key)
				}
				table.Append([]string{niceKey, strconv.Itoa(counterMap[key])})
				total += counterMap[key]
			}
			table.Append([]string{"Total:", strconv.Itoa(total)})
			table.Render()
		}
	}

	return nil
}
