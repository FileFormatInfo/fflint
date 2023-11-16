package main

import (
	"fmt"
	"os"

	"github.com/FileFormatInfo/fflint/internal/command"
	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/spf13/cobra"
)

var (
	version = "0.0.0"
	commit  = "local"
	date    = "local"
	builtBy = "unknown"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:           "fflint",
		Short:         "A linter to make sure your files are valid",
		Long:          `See [www.fflint.org](https://www.fflint.org/) for detailed instructions`,
		Version:       version,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	shared.AddCommon(rootCmd)

	command.AddAllCommands(rootCmd)
	command.AddVersionCommand(rootCmd, command.VersionInfo{Commit: commit, Version: version, LastMod: date, Builder: builtBy})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err.Error())
		os.Exit(1)
	}
}
