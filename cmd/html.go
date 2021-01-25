package cmd

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "html",
	Short: "test html files",
	Long:  `Validate that your html files are valid`,
	RunE:  makeFileCommand(htmlCheck),
}

func init() {
	rootCmd.AddCommand(htmlCmd)
}

func htmlCheck(f *FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.recordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	parseErr := validateHTML(bytes.NewReader(data))

	if parseErr != nil {
		f.recordResult("htmlParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}
}

func validateHTML(r *bytes.Reader) error {
	d := xml.NewDecoder(r)

	// Configure the decoder for HTML; leave off strict and autoclose for XHTML
	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity
	for {
		_, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func hasErrorNodes(node *html.Node) bool {
	if debug {
		fmt.Fprintf(os.Stderr, "Node type=%d text=%d doc=%d\n", node.Type, html.TextNode, html.DocumentNode)
	}

	if node.Type == html.ErrorNode {
		return true
	}

	if node.DataAtom == 0 {
		return true
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if hasErrorNodes(c) {
			return true
		}
	}
	return false
}
