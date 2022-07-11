package main

import (
	"fmt"
	"os"

	"github.com/fileformat/badger/internal/command"
	"github.com/fileformat/badger/internal/shared"
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
		Use:           "badger",
		Short:         "Badgers you if your file formats are invalid",
		Long:          `See [www.badger.sh](https://www.badger.sh/) for detailed instructions`,
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
