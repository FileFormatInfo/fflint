package command

import (
	"fmt"
	"os"

	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	yamlSorted     bool
	yamlStringKeys bool
)
var yamlCmd = &cobra.Command{
	Args:    cobra.MinimumNArgs(1),
	Use:     "yaml [options] files...",
	Short:   "Validate YAML files",
	Long:    `Check that your YAML files are valid`,
	PreRunE: yamlPrepare,
	RunE:    shared.MakeFileCommand(yamlCheck),
}

func AddYamlCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(yamlCmd)

	yamlCmd.Flags().BoolVar(&yamlSorted, "sorted", false, "Keys need to be in alphabetical order")
	yamlCmd.Flags().BoolVar(&yamlStringKeys, "stringkeys", true, "Keys need to be strings")
	//LATER: schema
	//LATER: string keys
}

func yamlCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	var yamlNode yaml.Node

	parseErr := yaml.Unmarshal(data, &yamlNode)

	if parseErr != nil {
		f.RecordResult("yamlParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}

	if yamlSorted || yamlStringKeys {
		sortErr := yamlWalker(&yamlNode, yamlSorted)
		f.RecordResult("yamlSorted", sortErr == nil, map[string]interface{}{
			"err": shared.ErrString(sortErr),
		})
	}
}

func yamlWalker(parent *yaml.Node, sorted bool) error {

	if parent.Kind == yaml.DocumentNode {
		return yamlWalker(parent.Content[0], sorted)
	}

	if parent.Kind != yaml.MappingNode {
		return nil
	}

	if len(parent.Content) > 0 {
		previousKey := ""
		for pos := 0; pos < len(parent.Content); pos += 2 {
			keyNode := parent.Content[pos]
			if keyNode.Tag != "!!str" {
				return fmt.Errorf("expected a string key but got %s on line %d, column %d", keyNode.Tag, keyNode.Line, keyNode.Column)
			}

			var key string
			keyDecodeErr := keyNode.Decode(&key)
			if keyDecodeErr != nil {
				return keyDecodeErr
			}

			if sorted {
				if shared.Debug {
					fmt.Fprintf(os.Stderr, "DEBUG: comparing %s < %s\n", previousKey, key)
				}
				if previousKey > key {
					return fmt.Errorf("keys out of order '%s' > '%s' (key '%s' at line %d, column %d", previousKey, key, key, keyNode.Line, keyNode.Column)
				}
				previousKey = key
			}

			valueNode := parent.Content[pos+1]
			if valueNode.Kind == yaml.MappingNode {
				childErr := yamlWalker(valueNode, sorted)
				if childErr != nil {
					return childErr
				}
			}
		}
	}

	return nil
}

func yamlPrepare(cmd *cobra.Command, args []string) error {
	return nil
}
