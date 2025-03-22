package webapp

import (
	"net/http"
	"runtime"
	"time"
)

var COMMIT string
var LASTMOD string

type Status struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Commit    string `json:"commit"`
	LastMod   string `json:"lastmod"`
	Tech      string `json:"tech"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := Status{}

	status.Success = true
	status.Message = "OK"
	status.Timestamp = time.Now().UTC().Format(time.RFC3339)
	status.Commit = COMMIT
	status.LastMod = LASTMOD
	status.Tech = runtime.Version()

	handleJson(w, r, status)
}
