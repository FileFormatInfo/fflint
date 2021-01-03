package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var COMMIT string = "local"
var LASTMOD string = "1970-01-01T00:00:01-00:00"
var VERSION = "v0.0.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of badger that is installed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s (%s - %s)\n", VERSION, COMMIT, LASTMOD)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
