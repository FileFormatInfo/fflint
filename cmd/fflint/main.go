package main

import (
	"fmt"
	"os"

	"github.com/FileFormatInfo/fflint/internal/command"
	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/spf13/cobra"
)

var (
	VERSION = "0.0.0"
	COMMIT  = "local"
	LASTMOD = "local"
	BUILTBY = "unknown"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:           "fflint",
		Short:         "A linter to make sure your files are valid",
		Long:          `See [www.fflint.dev](https://www.fflint.dev/) for detailed instructions`,
		Version:       VERSION,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	shared.AddCommon(rootCmd)

	command.AddAllCommands(rootCmd)
	command.AddVersionCommand(rootCmd, command.VersionInfo{Commit: COMMIT, Version: VERSION, LastMod: LASTMOD, Builder: BUILTBY})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err.Error())
		os.Exit(1)
	}
}
