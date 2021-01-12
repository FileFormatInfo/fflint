package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

//type CheckFn func(FileContext) void

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

func makeFileCommand(checkFn func(FileContext)) func(cmd *cobra.Command, args []string) {

	return func(cmd *cobra.Command, args []string) {

		files, _ := expandGlobs(args)

		for _, f := range files {
			basicTests(f)

			checkFn(f)
		}
	}

}
