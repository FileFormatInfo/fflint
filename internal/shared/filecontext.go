package shared

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
	//text     string
	tests []TestResult
}

// IsDir if it is a directory
func (fc *FileContext) IsDir() bool {
	fi, statErr := fc.Stat()
	if statErr != nil {
		if Debug {
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
		if Debug {
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

func (fc *FileContext) Size() int64 {
	if len(fc.data) > 0 {
		return int64(len(fc.data))
	}
	fi, sizeErr := fc.Stat()
	if sizeErr != nil {
		fmt.Fprintf(os.Stderr, "ERROR: unable to get size of file '%s' (%s)", fc.FilePath, sizeErr)
		os.Exit(4)
	}

	return fi.Size()
}

// ReadFile ioutil.ReadFile, possibly cached
func (fc *FileContext) ReadFile() ([]byte, error) {
	if len(fc.data) > 0 {
		return fc.data, nil
	}
	return ioutil.ReadFile(fc.FilePath)
}

func (fc *FileContext) RecordResult(Code string, Success bool, Detail map[string]interface{}) {
	fc.tests = append(fc.tests, TestResult{
		Code, Success, Detail,
	})

	fc.outputResult(Code, Success, Detail)

	if failFast && !Success {
		os.Exit(1)
	}
}

func (fc *FileContext) outputResult(Code string, Success bool, Detail map[string]interface{}) {
	if showTests.String() == "none" {
		return
	}

	if Success && showTests.String() == "failing" {
		return
	}

	if OutputFormat.String() == "json" {
		testData := map[string]interface{}{
			"file":    fc.FilePath,
			"success": Success,
			"test":    Code,
		}

		if showDetail {
			testData["detail"] = Detail
		}
		fmt.Printf("%s\n", EncodeJSON(testData))
	} else if OutputFormat.String() == "text" {
		fmt.Printf("INFO: %s %s %s", IfThenElse(Success, "PASS", "FAIL"), Code, fc.FilePath)
		if showDetail && Detail != nil {
			fmt.Printf(" %s", EncodeJSON(Detail))
		}

		fmt.Printf("\n")
	}
}

func (fc *FileContext) Reset() {
	fc.tests = nil
}

func (fc *FileContext) Success() bool {

	for _, test := range fc.tests {
		if !test.Success {
			return false
		}
	}

	return true
}

func basicTests(fc *FileContext) {
	if fileSize.Exists() {
		fc.RecordResult("fileSize", fileSize.Check(uint64(fc.Size())), map[string]interface{}{
			"actualSize":  fc.Size(),
			"desiredSize": fileSize.String(),
		})
	}
}

func EncodeJSON(data interface{}) string {

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
