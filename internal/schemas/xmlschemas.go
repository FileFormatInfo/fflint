package schemas

import (
	"embed"
	"io"
	"io/fs"
)

//go:embed xml
var xmlFS embed.FS

func GetXmlSchema(schema string) ([]byte, error) {
	fs, fsErr := fs.Sub(xmlFS, "xml")
	if fsErr != nil {
		return nil, fsErr
	}

	f, openErr := fs.Open(schema)
	if openErr != nil {
		return nil, openErr
	}
	defer f.Close()

	retVal, readErr := io.ReadAll(f)
	if readErr != nil {
		return nil, readErr
	}
	return retVal, nil
}
