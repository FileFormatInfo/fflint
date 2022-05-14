package command

import (
	"bytes"
	"fmt"

	"github.com/adrg/frontmatter"
	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	fmReport    bool
	fmStrict    bool
	fmStrictSet map[string]bool
	fmSorted    bool
	fmRequired  []string
	fmOptional  []string
	fmForbidden []string
)
var frontmatterCmd = &cobra.Command{
	Args:    cobra.MinimumNArgs(1),
	Use:     "frontmatter",
	Short:   "Validate frontmatter",
	Long:    `Checks that the frontmatter in your files is valid`,
	PreRunE: frontmatterPrepare,
	RunE:    shared.MakeFileCommand(frontmatterCheck),
}

func AddFrontmatterCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(frontmatterCmd)

	frontmatterCmd.Flags().StringSliceVar(&fmRequired, "required", []string{}, "Required keys")
	frontmatterCmd.Flags().StringSliceVar(&fmOptional, "optional", []string{}, "Optional keys (only for --strict)")
	frontmatterCmd.Flags().StringSliceVar(&fmForbidden, "forbidden", []string{}, "Forbidden keys")
	frontmatterCmd.Flags().BoolVar(&fmStrict, "strict", false, "Strict (keys must be in --required or --optional)")
	frontmatterCmd.Flags().BoolVar(&fmSorted, "sorted", false, "Keys need to be in alphabetical order")
	//optional
	//content: required/optional/forbidden
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

	_, parseErr := frontmatter.Parse(bytes.NewReader(data), &yamlData)

	if parseErr != nil {
		f.RecordResult("frontmatterParse", false, map[string]interface{}{
			"error": parseErr,
		})
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
			f.RecordResult("frontmatterSortedParse", previousKey < currentKey, map[string]interface{}{
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
	return nil
}
