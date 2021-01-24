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

// IsDir if it is a directory
func (fc *FileContext) IsDir() bool {
	fi, statErr := fc.Stat()
	if statErr != nil {
		if debug {
			fmt.Fprintf(os.Stderr, "DEBUG: error doing stat on %s: %s\n", fc.FilePath, statErr.Error())
		}
		return false
	}
	if fi.IsDir() {
		return true
	}
	return false
}

// IsFile if it is a file
func (fc *FileContext) IsFile() bool {
	fi, statErr := fc.Stat()
	if statErr != nil {
		if debug {
			fmt.Fprintf(os.Stderr, "DEBUG: error doing stat on %s: %s\n", fc.FilePath, statErr.Error())
		}
		return false
	}
	if fi.IsDir() {
		return false
	}
	//LATER: other tests?
	return true
}

// Stat os.Stat, possibly cached
func (fc *FileContext) Stat() (os.FileInfo, error) {
	return os.Stat(fc.FilePath) // MAYBE: cache?
}

// ReadFile ioutil.ReadFile, possibly cached
func (fc *FileContext) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(fc.FilePath)
}

func (fc *FileContext) recordResult(Code string, Success bool, Detail map[string]interface{}) {
	fc.tests = append(fc.tests, TestResult{
		Code, Success, Detail,
	})

	if !showTests {
		return
	}

	if Success && !showPassing {
		return
	}

	if outputFormat == "json" {
		testData := map[string]interface{}{
			"file":    fc.FilePath,
			"success": Success,
			"test":    Code,
		}

		if showDetail {
			testData["detail"] = Detail
		}
		fmt.Printf("%s\n", encodeJSON(testData))
	} else {
		fmt.Printf("INFO: %s %s %s", IfThenElse(Success, "PASS", "FAIL"), Code, fc.FilePath)
		if showDetail && Detail != nil {
			fmt.Printf(" %s", encodeJSON(Detail))
		}

		fmt.Printf("\n")
	}
}

func (fc *FileContext) reset() {
	fc.tests = nil
}

func (fc *FileContext) success() bool {

	for _, test := range fc.tests {
		if !test.Success {
			return false
		}
	}

	return true
}

func basicTests(fc *FileContext) {
	fi, err := fc.Stat()
	if err != nil {
		fc.recordResult("stat", false, map[string]interface{}{"error": err})
		return
	}

	if fileSize.Exists() {
		fc.recordResult("fileSize", fileSize.Check(uint64(fi.Size())), map[string]interface{}{
			"actualSize":  fi.Size(),
			"desiredSize": fileSize.String(),
		})
	}
}

func encodeJSON(data interface{}) string {

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
