package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func readRealLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func unbase64() int {
	log.Printf("unbase64: %#v\n", os.Args)
	filespec := "./*.base64"
	if len(os.Args) > 1 {
		filespec = os.Args[1]
	}
	files, globErr := filepath.Glob(filespec)
	if globErr != nil || len(files) == 0 {
		return 1
	}
	fmt.Fprintf(os.Stderr, "DEBUG: %d files found\n", len(files))

	for _, filename := range files {
		newfilename := filename[:len(filename)-7]
		fmt.Fprintf(os.Stderr, "DEBUG: decoding %s to %s\n", filename, newfilename)
		lines, readErr := readRealLines(filename)
		if readErr != nil {
			fmt.Fprintf(os.Stderr, "ERROR: unable to read %s: %v\n", filename, readErr)
			return 1
		}
		b, decodeErr := base64.StdEncoding.DecodeString(strings.Join(lines, ""))
		if decodeErr != nil {
			fmt.Fprintf(os.Stderr, "ERROR: unable to decode base64 in %s: %v\n", filename, decodeErr)
			return 1
		}
		f, writeErr := os.Create(newfilename)
		if writeErr != nil {
			fmt.Fprintf(os.Stderr, "ERROR: unable to write to %s: %v\n", filename, writeErr)
			return 1
		}
		defer f.Close()

		f.Write(b)
	}

	return 0
}

var extraCmds = map[string]func() int{
	"unbase64": unbase64,
}

func TestMain(m *testing.M) {
	exitVal := testscript.RunMain(m, extraCmds)

	os.Exit(exitVal)
}

func TestFflint(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "../../testdata",
	})
}
