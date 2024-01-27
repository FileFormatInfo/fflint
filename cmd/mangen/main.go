package main

import (
	"fmt"
	"os"

	"github.com/FileFormatInfo/fflint/internal/command"
	"github.com/FileFormatInfo/fflint/internal/shared"
	mango "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "fflint",
		Short: "A linter to make sure your files are valid",
		Long:  `See [www.fflint.dev](https://www.fflint.dev/) for detailed instructions`,
	}

	shared.AddCommon(rootCmd)

	command.AddExtCommand(rootCmd)
	command.AddFrontmatterCommand(rootCmd)
	command.AddHtmlCommand(rootCmd)
	command.AddIcoCommand(rootCmd)
	command.AddJpegCommand(rootCmd)
	command.AddJsonCommand(rootCmd)
	command.AddMimeTypeCommand(rootCmd)
	command.AddPngCommand(rootCmd)
	command.AddSvgCommand(rootCmd)
	command.AddTextCommand(rootCmd)
	command.AddVersionCommand(rootCmd, command.VersionInfo{})
	command.AddXmlCommand(rootCmd)
	command.AddYamlCommand(rootCmd)

	manPage, mangoErr := mango.NewManPage(1, rootCmd)
	if mangoErr != nil {
		fmt.Fprintf(os.Stderr, "ERROR: unable to generate man page: %v", mangoErr)
		os.Exit(1)
	}

	_, _ = fmt.Fprint(os.Stdout, manPage.Build(roff.NewDocument()))

}
