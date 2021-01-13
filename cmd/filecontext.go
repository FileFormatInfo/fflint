package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// TestResult is results of a single test
type TestResult struct {
	Code    string
	Success bool
	Detail  map[string]interface{}
}

// FileContext data about a file being tested
type FileContext struct {
	FilePath string
	data     []byte
	text     string
	tests    []TestResult
}

// Stat os.Stat, possibly cached
func (f *FileContext) Stat() (os.FileInfo, error) {
	return os.Stat(f.FilePath) // MAYBE: cache?
}

// ReadFile ioutil.ReadFile, possibly cached
func (f *FileContext) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(f.FilePath)
}

func (f *FileContext) recordResult(Code string, Success bool, Detail map[string]interface{}) {
	f.tests = append(f.tests, TestResult{
		Code, Success, Detail,
	})

	if silent {
		return
	}

	if Success && !showPassing {
		return
	}

	if outputFormat == "json" {
		fmt.Printf("%s\n", encodeJSON(map[string]interface{}{
			"detail":  Detail,
			"file":    f.FilePath,
			"success": Success,
			"test":    Code,
		}))
	} else {
		fmt.Printf("INFO: %s %s %s", IfThenElse(Success, "PASS", "FAIL"), Code, f.FilePath)
		if verbose && Detail != nil {
			fmt.Printf("%s", encodeJSON(Detail))
		}

		fmt.Printf("\n")
	}
}

func (f *FileContext) reset() {
	f.tests = nil
}

func (f *FileContext) success() bool {

	for _, test := range f.tests {
		if !test.Success {
			return false
		}
	}

	return true
}

func basicTests(f *FileContext) {
	fi, err := f.Stat()
	if err != nil {
		f.recordResult("stat", false, map[string]interface{}{"error": err})
		return
	}

	if fileSize.Exists() {
		f.recordResult("fileSize", fileSize.Check(uint64(fi.Size())), map[string]interface{}{
			"actualSize":  fi.Size(),
			"desiredSize": fileSize.String(),
		})
	}
}

func encodeJSON(data map[string]interface{}) string {

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	jsonErr := enc.Encode(data)

	if jsonErr != nil {
		// can this happen?
		return jsonErr.Error()
	}

	return strings.TrimRight(buf.String(), "\n")
}

// IfThenElse is a substitute for golang missing a ternary operator
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
