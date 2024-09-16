package command

import (
	"bytes"
	"encoding/xml"

	//"fmt"
	"io"
	//"os"

	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/spf13/cobra"
	//"golang.org/x/net/html"
)

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "html [options] files...",
	Short: "Validate HTML files",
	Long:  `Check HTML files for errors`, //LATER
	RunE:  shared.MakeFileCommand(htmlCheck),
}

func AddHtmlCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(htmlCmd)

	//LATER: no inline script
	//LATER: no inline styles
	//LATER: list,of,allowed,tags (or * or ones with html atoms?)
	//LATER: list of forbidden tags

}

func htmlCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	parseErr := validateHTML(bytes.NewReader(data))

	if parseErr != nil {
		f.RecordResult("htmlParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}
}

func validateHTML(r *bytes.Reader) error {
	d := xml.NewDecoder(r)

	path := []string{}

	//LATER: alternate parser [tdewolff/parse](https://github.com/tdewolff/parse)

	// Configure the decoder for HTML; leave off strict and autoclose for XHTML
	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity
	for {
		theToken, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			//fmt.Printf("err %T path=%v\n", err, path)
			return err
		}
		switch typedToken := theToken.(type) {
		case xml.StartElement:
			path = append(path, typedToken.Name.Local)
		case xml.EndElement:
			path = path[:len(path)-1]
		default:
			// ignore
		}

	}
}

/*
func hasErrorNodes(node *html.Node) bool {
	if shared.Debug {
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
*/
