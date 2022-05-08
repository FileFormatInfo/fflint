package command

import (
	"encoding/json"
	"fmt"

	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
)

var vi VersionInfo

type VersionInfo struct {
	Commit  string `json:"commit"`
	LastMod string `json:"lastmod"`
	Version string `json:"version"`
	Builder string `json:"builder"`
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Args:  cobra.NoArgs,
	Use:   "version",
	Short: "Prints badger version information",
	Run: func(cmd *cobra.Command, args []string) {
		if shared.OutputFormat == "json" {
			versionJSON, _ := json.Marshal(vi)
			fmt.Println(string(versionJSON))
		} else {
			if shared.Debug {
				fmt.Printf("Badger\n\tVersion: %s\n\tCommit: %s\n\tDate: %s\n\tBuilder: %s\n", vi.Version, vi.Commit, vi.LastMod, vi.Builder)
			} else {
				fmt.Printf("Badger v%s (%s)\n", vi.Version, vi.LastMod)
			}
		}
	},
}

func AddVersionCommand(rootCmd *cobra.Command, versionInfo VersionInfo) {
	rootCmd.AddCommand(versionCmd)
	vi = versionInfo
}
