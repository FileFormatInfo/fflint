package shared

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/mattn/go-isatty"

	"github.com/bmatcuk/doublestar/v3"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type globFn = func(args []string) ([]FileContext, error)

var globFunctions = map[string]globFn{
	"":           doublestarExpander,
	"doublestar": doublestarExpander,
	"golang":     golangExpander,
	"none":       noExpander,
}

// Globber is an optional min/max pair
type Globber struct {
	value string
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
	return "Glob algorithm"
}

func homedirExpand(arg string) string {
	expanded, err := homedir.Expand(arg)
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
		argfiles, _ := doublestar.Glob(homedirExpand(arg))
		for _, argfile := range argfiles {

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

func noExpander(args []string) ([]FileContext, error) {

	fcs := []FileContext{}

	for _, arg := range args {
		fc := FileContext{
			FilePath: homedirExpand(arg),
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

	return fcs, nil
}

func golangExpander(args []string) ([]FileContext, error) {

	fcs := []FileContext{}

	for _, arg := range args {
		argfiles, _ := filepath.Glob(homedirExpand(arg))
		for _, argfile := range argfiles {

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

func MakeFileCommand(checkFn func(*FileContext)) func(cmd *cobra.Command, args []string) error {

	return func(cmd *cobra.Command, args []string) error {

		total := 0
		bad := 0
		good := 0

		if Debug {
			fmt.Fprintf(os.Stderr, "DEBUG: %d args\n", len(args))
		}

		if len(args) == 1 {
			if args[0] == "-" && !isatty.IsTerminal(os.Stdin.Fd()) {
				scanner := bufio.NewScanner(os.Stdin)
				args = args[:0]
				for scanner.Scan() {
					line := scanner.Text()
					args = append(args, line)
				}
				if Debug {
					fmt.Fprintf(os.Stderr, "DEBUG: %d lines read from stdin\n", len(args))
				}
			}
			//LATER: handle @file
		}
		fcs, _ := globFunctions[globber.String()](args)
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
			if showFiles {
				if showPassing || !success {
					if OutputFormat == "json" {
						fileData := map[string]interface{}{
							"file":    fc.FilePath,
							"success": success,
						}

						fmt.Printf("%s\n", EncodeJSON(fileData))
					} else {
						fmt.Printf("INFO: %s %s", IfThenElse(success, "PASS", "FAIL"), fc.FilePath)

						fmt.Printf("\n")
					}

				}
			}
		}
		ProgressUpdate(good, bad, FileContext{})

		ProgressEnd()

		if showTotal {
			if OutputFormat == "json" {
				fmt.Printf("%s\n", EncodeJSON(map[string]interface{}{
					"total": total,
					"good":  good,
					"bad":   bad,
				}))
			} else {
				fmt.Printf("INFO: %d files tested, %d good, %d bad\n", total, good, bad)
			}
		}
		return nil
	}

}
