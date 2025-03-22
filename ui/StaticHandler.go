package ui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/FileFormatInfo/fflint/internal/common"
)

//go:embed static
var embeddedFiles embed.FS
var StaticHandler = initStaticHandler()

func initStaticHandler() http.Handler {

	fsys, err := fs.Sub(embeddedFiles, "static")
	if err != nil {
		common.Logger.Error("unable to create static file system", "error", err)
		panic(err)
	}

	return http.FileServer(http.FS(fsys))
}
