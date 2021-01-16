package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
		if debug {
			fmt.Fprintf(os.Stderr, "DEBUG: error while expanding homedir %s\n", err.Error())
		}
		return arg
	}
	return expanded
}

func doublestarExpander(args []string) ([]FileContext, error) {
	files := []FileContext{}

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

			files = append(files, fc)
		}
	}

	return files, nil
}

func noExpander(args []string) ([]FileContext, error) {

	files := []FileContext{}

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

		files = append(files, fc)
	}

	return files, nil
}

func golangExpander(args []string) ([]FileContext, error) {

	files := []FileContext{}

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

			files = append(files, fc)
		}
	}

	return files, nil
}

func makeFileCommand(checkFn func(*FileContext)) func(cmd *cobra.Command, args []string) error {

	return func(cmd *cobra.Command, args []string) error {

		total := 0
		bad := 0
		good := 0

		if debug {
			fmt.Fprintf(os.Stderr, "DEBUG: %d args\n", len(args))
		}
		files, _ := globFunctions[globber.String()](args)
		if debug {
			fmt.Fprintf(os.Stderr, "DEBUG: %d files after arg expansion\n", len(files))
		}

		ProgressStart(len(files))

		for _, fc := range files {
			basicTests(&fc)

			checkFn(&fc)

			total++
			success := fc.success()
			if success {
				good++
			} else {
				bad++
			}
			if showFiles {
				if showPassing || !success {
					if outputFormat == "json" {
						fileData := map[string]interface{}{
							"file":    fc.FilePath,
							"success": success,
						}

						fmt.Printf("%s\n", encodeJSON(fileData))
					} else {
						fmt.Printf("INFO: %s %s", IfThenElse(success, "PASS", "FAIL"), fc.FilePath)

						fmt.Printf("\n")
					}

				}
			}
			ProgressUpdate(success)
		}

		ProgressEnd()

		if showTotal {
			if outputFormat == "json" {
				fmt.Printf("%s\n", encodeJSON(map[string]interface{}{
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
