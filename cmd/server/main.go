package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/FileFormatInfo/fflint/internal/common"
	"github.com/FileFormatInfo/fflint/internal/webapp"
	"github.com/FileFormatInfo/fflint/ui"
)

func main() {

	var listenPort, portErr = strconv.Atoi(os.Getenv("PORT"))
	if portErr != nil {
		listenPort = 4000
	}
	var listenAddress = os.Getenv("ADDRESS")

	http.HandleFunc("/status.json", webapp.StatusHandler)
	http.HandleFunc("/{$}", webapp.RootHandler)
	http.HandleFunc("/robots.txt", ui.StaticHandler.ServeHTTP)
	http.HandleFunc("/favicon.ico", ui.StaticHandler.ServeHTTP)
	http.HandleFunc("/favicon.svg", ui.StaticHandler.ServeHTTP)
	//http.HandleFunc("/images/", ui.StaticHandler.ServeHTTP)

	http.HandleFunc("GET /language.html", webapp.LanguageGetHandler)
	http.HandleFunc("POST /language.html", webapp.LanguagePostHandler)

	contactUrl := os.Getenv("CONTACT_URL")
	if contactUrl == "" {
		contactUrl = "https://github.com/FileFormatInfo/fflint/issues"
	}
	http.HandleFunc("/contact.html", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, contactUrl, http.StatusFound)
	})

	http.HandleFunc("/", webapp.NotFoundHandler)

	err := http.ListenAndServe(listenAddress+":"+strconv.Itoa(listenPort), nil)
	if err != nil {
		common.Logger.Error("unable to listen", "address", listenAddress, "port", listenPort, "error", err)
	}
}
