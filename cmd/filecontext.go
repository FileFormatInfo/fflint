package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

func (f *FileContext) recordResult(Code string, Success bool, Detail map[string]interface{}) {
	f.tests = append(f.tests, TestResult{
		Code, Success, Detail,
	})

	if Success && !showPassing {
		return
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	jsonErr := enc.Encode(Detail)

	fmt.Printf("INFO: %s %s %s", IfThenElse(Success, "PASS", "FAIL"), Code, f.FilePath)

	if verbose {
		fmt.Printf("%s", IfThenElse(jsonErr != nil, jsonErr, strings.TrimRight(buf.String(), "\n")))
	}

	fmt.Printf("\n")
}

func expandGlobs(args []string) ([]FileContext, error) {

	if debug {
		fmt.Printf("DEBUG: %d args\n", len(args))
	}

	files := []FileContext{}

	for _, arg := range args {
		argfiles, _ := filepath.Glob(arg)
		for _, argfile := range argfiles {
			files = append(files, FileContext{
				FilePath: argfile,
			})
		}
	}

	if debug {
		fmt.Printf("DEBUG: %d files after arg expansion\n", len(files))
	}

	return files, nil
}

func basicTests(f FileContext) {
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

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
