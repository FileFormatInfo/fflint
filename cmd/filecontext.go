package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// TestResult: results of a single test
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

func (f *FileContext) RecordResult(Code string, Success bool, Detail map[string]interface{}) {
	f.tests = append(f.tests, TestResult{
		Code, Success, Detail,
	})

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	jsonErr := enc.Encode(Detail)

	fmt.Printf("INFO: %s: %s %s %s\n", f.FilePath, Code, IfThenElse(Success, "PASS", "FAIL"), IfThenElse(jsonErr != nil, jsonErr, buf.String()))
}

func expandGlobs(args []string) ([]FileContext, error) {
	files := []FileContext{}

	for _, arg := range args {
		argfiles, _ := filepath.Glob(arg)
		for _, argfile := range argfiles {
			files = append(files, FileContext{
				FilePath: argfile,
			})
		}
	}

	return files, nil
}

func basicTests(f FileContext) {
	fi, err := f.Stat()
	if err != nil {
		f.RecordResult("stat", false, map[string]interface{}{"error": err})
		return
	}

	f.RecordResult("minSize", fi.Size() >= minSize, map[string]interface{}{"size": fi.Size()})
	f.RecordResult("maxSize", fi.Size() <= maxSize, map[string]interface{}{"size": fi.Size()})
	if fileSize.Exists() {
		f.RecordResult("fileSize", fileSize.Check(uint64(fi.Size())), map[string]interface{}{
			"actualSize":  fi.Size(),
			"desiredSize": fileSize.String(),
		})
	}
}

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
