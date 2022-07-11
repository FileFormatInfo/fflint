package command

import "github.com/spf13/cobra"

func AddAllCommands(rootCmd *cobra.Command) {

	AddExtCommand(rootCmd)
	AddFrontmatterCommand(rootCmd)
	AddHtmlCommand(rootCmd)
	AddIcoCommand(rootCmd)
	AddJpegCommand(rootCmd)
	AddJsonCommand(rootCmd)
	AddMimeTypeCommand(rootCmd)
	AddPngCommand(rootCmd)
	AddSvgCommand(rootCmd)
	AddTextCommand(rootCmd)
	AddXmlCommand(rootCmd)
	AddYamlCommand(rootCmd)
}
