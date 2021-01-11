package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// COMMIT will be filled in a build time
var COMMIT string = "local"

// LASTMOD will be filled in a build time
var LASTMOD string = "1970-01-01T00:00:01-00:00"

// VERSION will be filled in a build time
var VERSION = "0.0.0"

type versionOutput struct {
	Commit  string `json:"commit"`
	LastMod string `json:"lastmod"`
	Version string `json:"version"`
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of badger that is installed",
	Run: func(cmd *cobra.Command, args []string) {
		if outputFormat == "json" {
			versionData := &versionOutput{
				Commit:  COMMIT,
				LastMod: LASTMOD,
				Version: VERSION,
			}
			versionJson, _ := json.Marshal(versionData)
			fmt.Println(string(versionJson))
		} else {
			fmt.Printf("Badger v%s (%s - %s)\n", VERSION, COMMIT, LASTMOD)
		}
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
