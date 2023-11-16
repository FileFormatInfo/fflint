package command

import (
	"bytes"
	"fmt"
	"os"

	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/spf13/cobra"
	"github.com/zyxar/image2ascii/ico"
)

var (
	icoStrict    bool
	icoStrictSet map[uint]bool
	icoRequired  []uint
	icoOptional  []uint
)

// icoCmd represents the ico command
var icoCmd = &cobra.Command{
	Args:    cobra.MinimumNArgs(1),
	Use:     "ico [options] files...",
	Short:   "Validate icons",
	Long:    `Check that your icons (.ico files) are valid`,
	PreRunE: icoPrepare,
	RunE:    shared.MakeFileCommand(icoCheck),
}

func AddIcoCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(icoCmd)

	icoCmd.Flags().UintSliceVar(&icoRequired, "required", []uint{16, 32}, "Required sizes")
	icoCmd.Flags().UintSliceVar(&icoOptional, "optional", []uint{48, 64, 96, 128}, "Optional sizes (only for `--strict`)")
	icoCmd.Flags().BoolVar(&icoStrict, "strict", true, "Strict (sizes must be in `--required` or `--optional`)")
	//LATER: allowPng
	//LATER: allowSvg
}

func uintContains(haystack []uint, needle uint) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}

func icoCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	images, parseErr := ico.DecodeAll(bytes.NewReader(data))

	if parseErr != nil {
		f.RecordResult("icoParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}

	sizes := []uint{}

	for index, image := range images {
		b := image.Bounds()
		width := b.Max.X - b.Min.X
		height := b.Max.Y - b.Min.Y
		if width < 0 || height < 0 {
			f.RecordResult("icoSize", false, map[string]interface{}{
				"bounds": b,
				"error":  "invalid image size",
				"height": height,
				"index":  index,
				"width":  width,
			})
			continue
		}
		if width != height {
			f.RecordResult("icoSquare", false, map[string]interface{}{
				"error":  "rectangular image",
				"height": height,
				"index":  index,
				"width":  width,
			})
			continue
		}
		sizes = append(sizes, uint(width))
	}
	if shared.Debug {
		fmt.Fprintf(os.Stderr, "DEBUG: icon sizes: %v\n", sizes)
	}

	if len(icoRequired) > 0 {
		for _, key := range icoRequired {
			f.RecordResult("icoRequired", uintContains(sizes, key), map[string]interface{}{
				"key": key,
			})
		}
	}

	if icoStrict {
		for _, key := range sizes {
			_, ok := icoStrictSet[key]
			f.RecordResult("icoStrict", ok, map[string]interface{}{
				"size": key,
			})
		}
	}
}

func icoPrepare(cmd *cobra.Command, args []string) error {
	if icoStrict {
		icoStrictSet = make(map[uint]bool)
		for _, key := range icoRequired {
			icoStrictSet[key] = true
		}
		for _, key := range icoOptional {
			icoStrictSet[key] = true
		}
	}
	return nil
}
