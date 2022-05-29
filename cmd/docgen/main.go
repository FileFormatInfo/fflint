package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/fileformat/badger/internal/command"
	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/pflag"
)

type cmdOption struct {
	Name         string
	Shorthand    string `yaml:",omitempty"`
	DefaultValue string `yaml:"default_value,omitempty"`
	Usage        string `yaml:",omitempty"`
}

type cmdDoc struct {
	Name             string
	Synopsis         string      `yaml:",omitempty"`
	Description      string      `yaml:",omitempty"`
	Usage            string      `yaml:",omitempty"`
	Options          []cmdOption `yaml:",omitempty"`
	InheritedOptions []cmdOption `yaml:"inherited_options,omitempty"`
	Example          string      `yaml:",omitempty"`
	SeeAlso          []string    `yaml:"see_also,omitempty"`
}

type indexEntry struct {
	Name        string
	Description string
}

func getFirstWord(s string) string {
	space := strings.IndexByte(s, ' ')
	if space == -1 {
		return s
	}
	return s[:space]
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "badger",
		Short: "Badgers you if your file formats are invalid",
		Long:  `See [www.badger.sh](https://www.badger.sh/) for detailed instructions`,
	}

	shared.AddCommon(rootCmd)

	command.AddExtCommand(rootCmd)
	command.AddFrontmatterCommand(rootCmd)
	command.AddHtmlCommand(rootCmd)
	command.AddIcoCommand(rootCmd)
	command.AddJpegCommand(rootCmd)
	command.AddJsonCommand(rootCmd)
	command.AddMimeTypeCommand(rootCmd)
	command.AddPngCommand(rootCmd)
	command.AddSvgCommand(rootCmd)
	command.AddTextCommand(rootCmd)
	command.AddVersionCommand(rootCmd, command.VersionInfo{})
	command.AddXmlCommand(rootCmd)
	command.AddYamlCommand(rootCmd)

	indexEntries := make([]indexEntry, 0)
	noteLine := "{% comment %}NOTE: this file is auto-generated by bin/docgen.sh.  Manual edits will be overwritten!{% endcomment -%}\n"
	for _, c := range rootCmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}

		cmdName := getFirstWord(c.Use)

		fmt.Fprintf(os.Stderr, "DEBUG: generating for %s\n", cmdName)

		var buf bytes.Buffer
		yamlErr := doc.GenYamlCustom(c, &buf, func(s string) string { return s })
		if yamlErr != nil {
			panic(yamlErr)
		}

		f, openErr := os.Create(fmt.Sprintf("./docs/_commands/%s.md", cmdName))
		if openErr != nil {
			panic(openErr)
		}
		w := bufio.NewWriter(f)
		w.WriteString("---\n")
		w.WriteString(fmt.Sprintf("h1: The %s Command\n", cmdName))
		w.WriteString(fmt.Sprintf("title: '%s: %s - Badger'\n", cmdName, c.Short))
		w.WriteString(buf.String())
		w.WriteString("---\n")
		w.WriteString(noteLine)
		w.WriteString("{% include command.html %}\n")
		w.Flush()
		f.Close()

		indexEntries = append(indexEntries, indexEntry{Name: cmdName, Description: c.Short})
	}

	f, openErr := os.Create(fmt.Sprintf("./docs/commands/index.md"))
	if openErr != nil {
		panic(openErr)
	}
	w := bufio.NewWriter(f)
	w.WriteString("---\n")
	w.WriteString("title: Available Commands\n")
	w.WriteString("---\n")
	w.WriteString(noteLine)
	w.WriteString("\n<table class=\"table table-bordered table-striped\">")
	for _, ie := range indexEntries {
		//LATER: make this into a table
		w.WriteString(fmt.Sprintf("<tr>\n<td><a href=\"%s.html\">%s</a></td>\n<td>%s</td>\n</tr>\n", ie.Name, ie.Name, ie.Description))
	}
	w.WriteString("</table>\n")
	w.Flush()
	f.Close()

	commmonFile, cfOpenErr := os.Create("./docs/_data/common_flags.yaml")
	if cfOpenErr != nil {
		panic(cfOpenErr)
	}
	rootCmd.PersistentFlags().VisitAll(func(rootFlag *pflag.Flag) {
		fmt.Fprintf(commmonFile, "- name: %s\n  default_value: \"%s\"\n  type: %s\n  usage: \"%s\"\n", rootFlag.Name, rootFlag.DefValue, rootFlag.Value.Type(), rootFlag.Usage)
	})
	f.Close()

	//LATER: generate list of common flags
}
