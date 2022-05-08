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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// textCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
