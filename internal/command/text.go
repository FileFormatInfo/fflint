package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Validate plain text files",
	Long:  `Checks that plain text files really are plain text, have the correct line endings and more`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("text called")
	},
}

func AddTextCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(textCmd)

	//LATER: UTF-8
	//LATER: command: ASCII
	//LATER: tabs (bool)
	//LATER: bom
	/*
		LATER: many of these would be good for other text-based formats
		- charset:ascii|utf-8
		- trailing-newline: on/off/any/only
		- newline format: cr/crlf/lf/any (or dos/unix/mac?)
		- indent: tab/spaces/any
		- contains/doesnotcontain: specific text (license declaration/etc)
		- unicode: list of unicode character ranges allowed
	*/
}
