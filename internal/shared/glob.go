package shared

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/mattn/go-isatty"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type globFn = func(args []string) ([]FileContext, error)

//LATER: regex-based globber
//LATER: glob with https://github.com/gobwas/glob
var globFunctions = map[string]globFn{
	"":           doublestarExpander,
	"doublestar": doublestarExpander,
	"golang":     golangExpander,
	"none":       noExpander,
}

// Globber is an optional min/max pair
type Globber struct {
	value string `default:"doublestar"`
}

func (g *Globber) String() string {
	return g.value
}

// Set the glob algorithm
func (g *Globber) Set(newValue string) error {

	if globFunctions[newValue] == nil {
		return fmt.Errorf("Invalid glob algorithm '%s'", newValue)
	}
	g.value = newValue
	return nil
}

// Type is a description of range
func (g *Globber) Type() string {
	return "Globber"
}

func homedirExpand(arg string) string {
	expanded, err := homedir.Expand(arg) //LATER: switch to os.UserHomeDir()?
	if err != nil {
		if Debug {
			fmt.Fprintf(os.Stderr, "DEBUG: error while expanding homedir %s\n", err.Error())
		}
		return arg
	}
	return expanded
}

func doublestarExpander(args []string) ([]FileContext, error) {
	fcs := []FileContext{}

	for _, arg := range args {
		basepath, pattern := doublestar.SplitPattern(homedirExpand(arg))
		if Debug {
			fmt.Fprintf(os.Stderr, "DEBUG: doublestar expanding %s at %s\n", pattern, basepath)
		}
		fsys := os.DirFS(basepath)
		argfiles, dsErr := doublestar.Glob(fsys, pattern)
		if dsErr != nil {
			return nil, fmt.Errorf("Unable to expand %s at %s (doublestarExpander %w)", pattern, basepath, dsErr)
		}
		for _, argfile := range argfiles {

			if ignoreDotFiles && argfile[0] == '.' {
				continue
			}

			fullpath := filepath.Join(basepath, argfile)

			fc := FileContext{
				FilePath: fullpath,
			}

			fi, statErr := fc.Stat()
			if statErr != nil {
				return nil, fmt.Errorf("Unable to stat %s (doublestarExpander, %w)", arg, statErr)
			}

			if fi.IsDir() {
				if Debug {
					fmt.Fprintf(os.Stderr, "WARNING: doublestar returned a directory: %s\n", fullpath)
				}
				continue
			}
			ProgressCount()
			fcs = append(fcs, fc)
		}
	}

	return fcs, nil
}

func noExpander(args []string) ([]FileContext, error) {

	fcs := []FileContext{}

	for _, arg := range args {
		fc := FileContext{
			FilePath: homedirExpand(arg),
		}

		fi, statErr := fc.Stat()
		if statErr != nil {
			return nil, fmt.Errorf("Unable to stat %s (noExpander, %w)", arg, statErr)
		}
		if fi.IsDir() {
			//LATER: or recurse?
			continue
		}

		ProgressCount()
		fcs = append(fcs, fc)
	}

	return fcs, nil
}

func golangExpander(args []string) ([]FileContext, error) {

	fcs := []FileContext{}

	for _, arg := range args {
		argfiles, globErr := filepath.Glob(homedirExpand(arg))
		if globErr != nil {
			return nil, fmt.Errorf("Unable to glob %s (golangExpander, %w)", arg, globErr)
		}
		for _, argfile := range argfiles {

			if ignoreDotFiles && argfile[0] == '.' {
				continue
			}

			fc := FileContext{
				FilePath: argfile,
			}

			fi, statErr := fc.Stat()
			if statErr != nil {
				//LATER
				continue
			}
			if fi.IsDir() {
				//LATER: or recurse?
				continue
			}

			ProgressCount()
			fcs = append(fcs, fc)
		}
	}

	return fcs, nil
}

/*
func loadIgnoreFile() (*gitignore.Gitignore, error) {
	ignorer, err := gitignore.CompileIgnoreFile(".gitignore")

	return ignorer, err
}
*/

func MakeFileCommand(checkFn func(*FileContext)) func(cmd *cobra.Command, args []string) error {

	return func(cmd *cobra.Command, args []string) error {

		total := 0
		bad := 0
		good := 0

		if Debug {
			fmt.Fprintf(os.Stderr, "DEBUG: %d args\n", len(args))
		}

		fcs := []FileContext{}
		useGlobber := true
		if len(args) == 1 && args[0] == "-" {
			useGlobber = false
			if isatty.IsTerminal(os.Stdin.Fd()) {
				return errors.New("Attempt to read stdin from terminal")
			}
			data, stdinReadErr := ioutil.ReadAll(os.Stdin) //LATER: does this work on Windows w/binary files
			if stdinReadErr != nil {
				return fmt.Errorf("unable to read stdin: %w", stdinReadErr)
			}
			fcs = append(fcs, FileContext{
				FilePath: "stdin",
				data:     data,
			})
		} else if len(args) == 1 && args[0] == "@-" {
			if isatty.IsTerminal(os.Stdin.Fd()) {
				return errors.New("Attempt to read stdin from terminal")
			}
			scanner := bufio.NewScanner(os.Stdin)
			args = args[:0]
			for scanner.Scan() {
				line := scanner.Text()
				args = append(args, line)
			}
			if Debug {
				fmt.Fprintf(os.Stderr, "DEBUG: %d lines read from stdin\n", len(args))
			}
			useGlobber = false
			fcs, _ = noExpander(args)
		}
		//LATER: handle @file, @-0 for names on stdin
		if useGlobber {
			fcs, _ = globFunctions[globber.String()](args)
		}
		if len(fcs) == 0 {
			return fmt.Errorf("No files to badger")
		}
		if Debug {
			fmt.Fprintf(os.Stderr, "DEBUG: %d files after arg expansion\n", len(fcs))
		}
		sort.Slice(fcs[:], func(i, j int) bool {
			return fcs[i].FilePath < fcs[j].FilePath
		})

		ProgressStart(fcs)

		for _, fc := range fcs {

			ProgressUpdate(good, bad, fc)

			basicTests(&fc)

			checkFn(&fc)

			total++
			success := fc.Success()
			if success {
				good++
			} else {
				bad++
			}
			if showFiles.String() != "none" {
				if !success || showFiles.String() == "all" {
					if OutputFormat.String() == "json" {
						fileData := map[string]interface{}{
							"file":    fc.FilePath,
							"success": success,
						}

						fmt.Printf("%s\n", EncodeJSON(fileData))
					} else if OutputFormat.String() == "text" {
						fmt.Printf("%s: %s\n", IfThenElse(success, "INFO", "ERROR"), fc.FilePath)
					}
				}
			} else if OutputFormat.String() == "filenames" {
				if !success {
					fmt.Printf("%s\n", fc.FilePath)
				}
			}
		}
		ProgressUpdate(good, bad, FileContext{})

		ProgressEnd()

		if showTotal {
			if OutputFormat.String() == "json" {
				fmt.Printf("%s\n", EncodeJSON(map[string]interface{}{
					"total": total,
					"good":  good,
					"bad":   bad,
				}))
			} else if OutputFormat.String() == "text" {
				fmt.Printf("INFO: %d files tested, %d good, %d bad\n", total, good, bad)
			}
		}
		if bad > 0 {
			return fmt.Errorf("%d errors found", bad)
		}
		return nil
	}
}
