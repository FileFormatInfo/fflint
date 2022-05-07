package main

import (
	"fmt"
	"os"

	"github.com/fileformat/badger/internal/command"
	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "badger",
		Short: "Badgers you if your file formats are invalid",
		Long:  `See [www.badger.sh](https://www.badger.sh/) for detailed instructions`,
	}

	shared.AddCommon(rootCmd)
	command.AddSvgCommand(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: unable to execute root command: %s\n", err.Error())
		os.Exit(1)
	}
}
