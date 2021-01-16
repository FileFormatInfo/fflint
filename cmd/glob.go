package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

type globFn = func(args []string) ([]FileContext, error)

var globFunctions = map[string]globFn{
	"golang": golangExpander,
	"":       golangExpander,
	"none":   noExpander,
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

func noExpander(args []string) ([]FileContext, error) {

	if debug {
		fmt.Printf("DEBUG: %d args\n", len(args))
	}

	files := []FileContext{}

	for _, arg := range args {
		fc := FileContext{
			FilePath: arg,
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

	if debug {
		fmt.Printf("DEBUG: %d args\n", len(args))
	}

	files := []FileContext{}

	for _, arg := range args {
		argfiles, _ := filepath.Glob(arg)
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

	if debug {
		fmt.Printf("DEBUG: %d files after arg expansion\n", len(files))
	}

	return files, nil
}

func makeFileCommand(checkFn func(*FileContext)) func(cmd *cobra.Command, args []string) error {

	return func(cmd *cobra.Command, args []string) error {

		total := 0
		bad := 0

		files, _ := globFunctions[globber.String()](args)

		ProgressStart(len(files))

		for _, fc := range files {
			basicTests(&fc)

			checkFn(&fc)

			total++
			if !fc.success() {
				bad++
			}
			ProgressUpdate(fc.success())
		}

		ProgressEnd()

		if debug {
			fmt.Printf("DEBUG: %d files, %d bad", total, bad)
		}
		return nil
	}

}
