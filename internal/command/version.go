package command

import (
	"encoding/json"
	"fmt"

	"github.com/FileFormatInfo/fflint/internal/shared"
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
	Short: "Prints fflint version information",
	Run: func(cmd *cobra.Command, args []string) {
		if shared.OutputFormat.String() == "json" {
			versionJSON, _ := json.Marshal(vi)
			fmt.Println(string(versionJSON))
		} else if shared.OutputFormat.String() == "text" {
			fmt.Printf("Badger v%s (%s)\n", vi.Version, vi.LastMod)
			if shared.Debug {
				fmt.Printf("\tCommit: %s\n\tBuilder: %s\n", vi.Commit, vi.Builder)
			}
		}
	},
}

func AddVersionCommand(rootCmd *cobra.Command, versionInfo VersionInfo) {
	rootCmd.AddCommand(versionCmd)
	vi = versionInfo
}
