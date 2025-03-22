package webapp

import (
	"encoding/json"
	"net/http"

	"github.com/FileFormatInfo/fflint/internal/common"
)

func handleJson(w http.ResponseWriter, r *http.Request, data any) {

	b, err := json.Marshal(data)
	if err != nil {
		common.Logger.Error("json.Marshal failed", "error", err, "data", data)
		b = []byte("{\"success\":false,\"err\":\"json.Marshal failed\"}")
	}

	var callback = r.FormValue("callback")
	if callback != "" {
		w.Header().Set("Content-Type", "application/javascript; charset=utf8")
		w.Write([]byte(callback + "("))
		w.Write(b)
		w.Write([]byte(");"))
	} else {
		//w.Header().Set("Content-Type", "application/json; charset=utf8")
		w.Header().Set("Content-Type", "text/plain; charset=utf8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		w.Header().Set("Access-Control-Max-Age", "604800") // 1 week
		w.Write(b)
	}
}
