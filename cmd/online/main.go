package main

import (
	"bufio"
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/fileformat/badger/internal/online"
	flag "github.com/spf13/pflag"
)

var (
	version  = "0.0.0"
	commit   = "local"
	date     = "local"
	builtBy  = "unknown"
	port     int
	bindHost string
	//go:embed assets
	embeddedFiles embed.FS
)

func main() {

	flag.IntVar(&port, "port", 4000, "port")
	flag.StringVar(&bindHost, "bind", "", "IP address to bind to (localhost to avoid MacOS popup)")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/", command_handler)
	mux.HandleFunc("/status.json", status_handler)

	fsys, err := fs.Sub(embeddedFiles, "assets")
	if err != nil {
		panic(err)
	}

	mux.Handle("/", http.FileServer(http.FS(fsys)))
	http.ListenAndServe(fmt.Sprintf("%s:%d", bindHost, port), mux)
}

type statusResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Commit    string `json:"commit"`
	LastMod   string `json:"lastmod"`
	Tech      string `json:"tech"`
	BuiltBy   string `json:"builtby"`
	Version   string `json:"version"`
	Getwd     string `json:"os.Getwd()"`
	//Hostname  string `json:"os.Hostname()"`
	//TempDir   string `json:"os.TempDir()"`
}

func status_handler(w http.ResponseWriter, r *http.Request) {

	var status = statusResponse{
		Success:   true,
		Message:   "OK",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Commit:    commit,
		LastMod:   date,
		Version:   version,
		BuiltBy:   builtBy,
		Tech:      runtime.Version(),
	}

	wd, err := os.Getwd()
	if err != nil {
		status.Getwd = "ERROR: " + err.Error()
	} else {
		status.Getwd = wd
	}

	online.WriteJsonp(w, r, status)
}

type errorResponse struct {
	Success   bool              `json:"success"`
	Code      string            `json:"code"`
	Message   string            `json:"message"`
	Timestamp string            `json:"timestamp"`
	Detail    map[string]string `json:"detail,omitempty"`
}

var MAX_UPLOAD_SIZE int64 = 10 * 1024 * 1024 //LATER: from param

type badgerResponse struct {
	Success   bool          `json:"success"`
	Code      string        `json:"code"`
	Message   string        `json:"message"`
	Timestamp string        `json:"timestamp"`
	Results   []interface{} `json:"results,omitempty"`
	Stderr    string        `json:"stderr"`
}

func command_handler(w http.ResponseWriter, r *http.Request) {

	commandPath := r.URL.Path[5:]
	if commandPath == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if strings.HasSuffix(commandPath, ".json") == false {
		online.WriteJsonp(w, r, errorResponse{
			Success:   false,
			Code:      "InvalidOutputFormat",
			Message:   "Only JSON is currently supported",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	if r.Method != "POST" {
		online.WriteJsonp(w, r, errorResponse{
			Success:   false,
			Code:      "InvalidMethod",
			Message:   "Only POST is currently supported",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	commandName := strings.TrimSuffix(commandPath, ".json")

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	parseErr := r.ParseMultipartForm(MAX_UPLOAD_SIZE)
	if parseErr != nil {
		online.WriteJsonp(w, r, errorResponse{
			Success:   false,
			Code:      "UnableToParseForm",
			Message:   parseErr.Error(),
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	_, fileHeader, fileErr := r.FormFile("file") //LATER: loop through multiple files
	if fileErr != nil {
		online.WriteJsonp(w, r, errorResponse{
			Success:   false,
			Code:      "NoFileUploaded",
			Message:   fileErr.Error(),
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	f, openErr := fileHeader.Open()
	if openErr != nil {
		online.WriteJsonp(w, r, errorResponse{
			Success:   false,
			Code:      "FileOpenError",
			Message:   openErr.Error(),
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	defer f.Close()

	data, readErr := ioutil.ReadAll(f)
	if readErr != nil {
		online.WriteJsonp(w, r, errorResponse{
			Success:   false,
			Code:      "FileReadError",
			Message:   readErr.Error(),
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	args := []string{
		commandName,
	}

	args = append(args, "--show-files=none")
	args = append(args, "--show-tests=all")
	args = append(args, "--show-totals=false")
	args = append(args, "--debug=true")
	args = append(args, "--progress=false")
	args = append(args, "--output=json")
	args = append(args, "-")

	cmd := exec.Command("./badger", args...)

	cmd.Stdin = bytes.NewReader(data)

	var out bytes.Buffer
	cmd.Stdout = &out

	var errbuf bytes.Buffer
	cmd.Stderr = &errbuf

	runErr := cmd.Run()

	var testResults []interface{}

	scanner := bufio.NewScanner(bytes.NewReader(out.Bytes()))
	for scanner.Scan() {
		var tmp map[string]interface{}
		line := scanner.Text()
		jsonErr := json.Unmarshal([]byte(line), &tmp)
		if jsonErr != nil {
			fmt.Fprintf(os.Stderr, "Error %v for %s\n", jsonErr, line)
		} else {
			testResults = append(testResults, tmp)
		}
	}

	message := "Success"
	if runErr != nil {
		message = runErr.Error()
	}

	online.WriteJsonp(w, r, badgerResponse{
		Success:   runErr == nil,
		Message:   message,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Results:   testResults,
		Stderr:    string(errbuf.Bytes()),
	})
}
