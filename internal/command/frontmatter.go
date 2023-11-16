package command

import (
	"bytes"
	"fmt"
	"os"

	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/adrg/frontmatter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	fmReport     bool
	fmStrict     bool
	fmStrictSet  map[string]bool
	fmSorted     bool
	fmRequired   []string
	fmOptional   []string
	fmForbidden  []string
	fmDelimiters []string
)
var frontmatterCmd = &cobra.Command{
	Args:    cobra.MinimumNArgs(1),
	Use:     "frontmatter [options] files...",
	Short:   "Validate frontmatter",
	Long:    `Checks that the frontmatter in your files is valid`,
	PreRunE: frontmatterPrepare,
	RunE:    shared.MakeFileCommand(frontmatterCheck),
}

func AddFrontmatterCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(frontmatterCmd)

	frontmatterCmd.Flags().StringSliceVar(&fmRequired, "required", []string{}, "Required keys")
	frontmatterCmd.Flags().StringSliceVar(&fmOptional, "optional", []string{}, "Optional keys (only for `--strict`)")
	frontmatterCmd.Flags().StringSliceVar(&fmForbidden, "forbidden", []string{}, "Forbidden keys")
	frontmatterCmd.Flags().BoolVar(&fmStrict, "strict", false, "Strict (keys must be in `--required` or `--optional`)")
	frontmatterCmd.Flags().BoolVar(&fmSorted, "sorted", false, "Keys need to be in alphabetical order")
	frontmatterCmd.Flags().StringSliceVar(&fmDelimiters, "delimiters", []string{}, "Custom delimiters (if other than `---`, `+++` and `;;;`)")

	//LATER: report
	//LATER: schema
}

func frontmatterCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	yamlData := make(map[interface{}]interface{})

	var formats []*frontmatter.Format

	if len(fmDelimiters) > 0 {
		var end = fmDelimiters[0]
		if len(fmDelimiters) > 1 {
			end = fmDelimiters[1]
		}
		formats = append(formats, frontmatter.NewFormat(fmDelimiters[0], end, yaml.Unmarshal))
		if shared.Debug {
			fmt.Fprintf(os.Stderr, "DEBUG: using custom delimiters '%s' and '%s'\n", fmDelimiters[0], end)
		}
	}

	_, parseErr := frontmatter.MustParse(bytes.NewReader(data), &yamlData, formats...)

	f.RecordResult("frontmatterParse", parseErr == nil, map[string]interface{}{
		"error": shared.ErrString(parseErr),
	})
	if parseErr != nil {
		return
	}

	if len(fmRequired) > 0 {
		for _, key := range fmRequired {
			_, ok := yamlData[key]
			f.RecordResult("frontmatterRequired", ok, map[string]interface{}{
				"key": key,
			})
		}
	}

	if len(fmForbidden) > 0 {
		for _, key := range fmForbidden {
			_, ok := yamlData[key]
			f.RecordResult("frontmatterForbidden", !ok, map[string]interface{}{
				"key": key,
			})
		}
	}

	if fmStrict {
		for key := range yamlData {
			keyStr, strErr := key.(string)
			if !strErr {
				f.RecordResult("frontmatterStrictParse", false, map[string]interface{}{
					"err": "key is not a string",
					"key": fmt.Sprintf("%v", key),
				})
				continue
			}
			_, ok := fmStrictSet[keyStr]
			f.RecordResult("frontmatterStrict", ok, map[string]interface{}{
				"key": keyStr,
			})
		}
	}

	if fmSorted {
		sortedData := yaml.MapSlice{}

		frontmatter.Parse(bytes.NewReader(data), &sortedData)
		previousKey := ""
		for _, item := range sortedData {
			currentKey, strErr := item.Key.(string)
			if !strErr {
				f.RecordResult("frontmatterSortedParse", false, map[string]interface{}{
					"err": "key is not a string",
					"key": fmt.Sprintf("%v", item.Key),
				})
				continue
			}
			f.RecordResult("frontmatterSorted", previousKey < currentKey, map[string]interface{}{
				"previous": previousKey,
				"current":  currentKey,
			})
			previousKey = currentKey
		}
	}
}

func frontmatterPrepare(cmd *cobra.Command, args []string) error {
	if fmStrict {
		fmStrictSet = make(map[string]bool)
		for _, key := range fmRequired {
			fmStrictSet[key] = true
		}
		for _, key := range fmOptional {
			fmStrictSet[key] = true
		}
	}

	if len(fmDelimiters) > 2 {
		fmt.Fprintf(os.Stderr, "ERROR: delimiter count must be <=2 (passed %d)", len(fmDelimiters))
		os.Exit(7)
	}
	return nil
}
