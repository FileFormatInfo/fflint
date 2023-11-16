package command

import (
	"bytes"

	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/antchfx/xmlquery"
	"github.com/spf13/cobra"
)

var (
	schema string
	//xsdSchema *xsd.Schema
)

// xmlCmd represents the xml command
var xmlCmd = &cobra.Command{
	Args:     cobra.MinimumNArgs(1),
	Use:      "xml [options] files...",
	Short:    "Validate XML files",
	Long:     `Checks that your XML files are valid`,
	PreRunE:  xmlInit,
	RunE:     shared.MakeFileCommand(xmlCheck),
	PostRunE: xmlCleanup,
}

func AddXmlCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(xmlCmd)
	xmlCmd.Flags().StringVar(&schema, "schema", "", "Schema (XSD) to use") //LATER: link to docs about embedded ones
}

func xmlCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	_, parseErr := xmlquery.Parse(bytes.NewReader(data))

	if parseErr != nil {
		f.RecordResult("xmlParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}

	/*
		doc, parseErr := libxml2.ParseString(string(data))
		if parseErr != nil {
			f.RecordResult("xmlParse", false, map[string]interface{}{
				"error": shared.ErrString(parseErr),
			})
			return
		}
		defer doc.Free()

		if xsdSchema != nil {
			xsdErr := xsdSchema.Validate(doc)
			f.RecordResult("xmlSchema", xsdErr == nil, map[string]interface{}{
				"error": shared.ErrString(xsdErr),
			})
		}
	*/
}

func xmlInit(cmd *cobra.Command, args []string) error {

	if schema == "" {
		return nil
	}
	/*
		data, getErr := schemas.GetXmlSchema(schema)
		if getErr != nil {
			return getErr
		}

		var parseErr error
		xsdSchema, parseErr = xsd.Parse(data)
		if parseErr != nil {
			return parseErr
		}
	*/
	return nil
}

func xmlCleanup(cmd *cobra.Command, args []string) error {
	/*
		if xsdSchema != nil {
			xsdSchema.Free()
		}
	*/
	return nil
}

//LATER: schema validation

//NO: wrappers around libxml2 :(
// https://github.com/lestrrat-go/libxml2 /xsd
// https://github.com/terminalstatic/go-xsd-validate
// https://github.com/krolaw/xsd

// or generate for specific xsd's: https://github.com/xuri/xgen
// https://github.com/droyo/go-xml
