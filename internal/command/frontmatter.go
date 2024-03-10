package command

import (
	"bytes"
	"fmt"
	"os"

	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/adrg/frontmatter"
	"github.com/spf13/cobra"
	yamlv2 "gopkg.in/yaml.v2"
	"gopkg.in/yaml.v3"
)

var (
	fmSchemaOptions shared.SchemaOptions
	fmStrict        bool
	fmStrictSet     map[string]bool
	fmSorted        bool
	fmRequired      []string
	fmOptional      []string
	fmForbidden     []string
	fmDelimiters    []string
)
var frontmatterCmd = &cobra.Command{
	Args:    cobra.MinimumNArgs(1),
	Use:     "frontmatter [options] files...",
	Short:   "Validate frontmatter",
	Long:    `Checks that the frontmatter in your files is valid`,
	PreRunE: frontmatterInit,
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

	fmSchemaOptions.AddFlags(frontmatterCmd)

	//LATER: report
}

func frontmatterCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

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

	yamlRawData := make(map[any]any)
	//LATER: maybe flag to require contents?
	_, parseErr := frontmatter.MustParse(bytes.NewReader(data), &yamlRawData, formats...)

	f.RecordResult("frontmatterParse", parseErr == nil, map[string]interface{}{
		"error": shared.ErrString(parseErr),
	})
	if parseErr != nil {
		return
	}

	/*
		yamlDataOrArray, stringKeysErr := shared.ToStringKeys(yamlRawData)
		f.RecordResult("frontmatterStringKeys", stringKeysErr == nil, map[string]interface{}{
			"error": stringKeysErr,
		})
	*/
	yamlDataOrArray := convert(yamlRawData)
	yamlData := yamlDataOrArray.(map[string]interface{})

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
			_, ok := fmStrictSet[key]
			f.RecordResult("frontmatterStrict", ok, map[string]interface{}{
				"key": key,
			})
		}
	}

	if fmSorted {
		sortedData := yamlv2.MapSlice{}

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

	fmSchemaOptions.Validate(f, yamlData)
}

func frontmatterInit(cmd *cobra.Command, args []string) error {
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

	schemaPrepErr := fmSchemaOptions.Prepare()
	if schemaPrepErr != nil {
		return schemaPrepErr
	}

	return nil
}

// from https://stackoverflow.com/a/40737676
func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case map[string]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}
