package shared

import (
	"fmt"
	"os"

	gitignore "github.com/sabhiram/go-gitignore"
)

type ignoreFn = func(fileName string) bool

var (
	DEFAULT_IGNORE_FILE = ".gitignore"
)

func loadIgnoreFile(fileName string) ignoreFn {

	if fileName == "" {
		if Debug {
			fmt.Fprintf(os.Stderr, "DEBUG: ignore file disabled\n")
		}
		return func(fileName string) bool { return false }
	}

	_, statErr := os.Stat(fileName)
	if statErr != nil {
		if fileName == DEFAULT_IGNORE_FILE { //LATER: preferrable to use pflags.Changed, but would need a global var for rootCmd...
			if Debug {
				fmt.Fprintf(os.Stderr, "DEBUG: skipping default ignore file %s\n", fileName)
			}
			return func(fileName string) bool { return false }
		}
		fmt.Fprintf(os.Stderr, "ERROR: unable to find ignore file '%s': %v\n", fileName, statErr)
		os.Exit(5)
	}

	ignorer, err := gitignore.CompileIgnoreFile(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: unable to compile ignore file %s: %v", fileName, err)
		os.Exit(6)
	}

	if Debug {
		fmt.Fprintf(os.Stderr, "DEBUG: loaded ignore file %s\n", fileName)
	}

	return func(target string) bool {
		matches, how := ignorer.MatchesPathHow(target)

		if Debug && matches {
			fmt.Fprintf(os.Stderr, "DEBUG: skipping %s (%s: line %d)\n", target, fileName, how.LineNo)
		}

		return matches
	}
}
