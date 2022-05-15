package shared

import (
	"encoding/xml"
	"errors"
	"io"

	"golang.org/x/net/html/charset"
)

type Namespaces struct {
	Default    string
	Additional map[string]string
}

func parseNamespace(root *xml.StartElement) (Namespaces, error) {
	retVal := Namespaces{Default: "", Additional: make(map[string]string)}

	for _, attr := range root.Attr {
		if attr.Name.Space == "" && attr.Name.Local == "xmlns" {
			retVal.Default = attr.Value
		}
		if attr.Name.Space == "xmlns" {
			retVal.Additional[attr.Name.Local] = attr.Value
		}
	}

	return retVal, nil
}

func getRoot(decoder *xml.Decoder) (*xml.StartElement, error) {
	for {
		token, err := decoder.RawToken()
		if token == nil && err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		retVal, ok := token.(xml.StartElement)
		if ok {
			return &retVal, nil
		}
	}
	return nil, errors.New("root element not found")
}

func GetNamespaces(reader io.Reader) (Namespaces, error) {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	token, err := getRoot(decoder)
	if err != nil {
		return Namespaces{}, err
	}
	return parseNamespace(token)
}
