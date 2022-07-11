package command

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/fileformat/badger/internal/shared"
	"github.com/spf13/cobra"
)

var (
	textUtf8    bool
	textAscii   bool
	textControl bool
	textNul     bool
)

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text [options] files...",
	Short: "Validate plain text files",
	Long:  `Checks that plain text files really are plain text, have the correct line endings and more`,
	RunE:  shared.MakeFileCommand(textCheck),
}

func AddTextCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(textCmd)

	textCmd.Flags().BoolVar(&textUtf8, "utf8", true, "Must be valid UTF-8")
	textCmd.Flags().BoolVar(&textAscii, "ascii", false, "Must be valid ASCII")
	textCmd.Flags().BoolVar(&textControl, "control", false, "Allow control characters (other than newlines and tab)")
	textCmd.Flags().BoolVar(&textNul, "nul", false, "Allow nul (0x00) bytes")
	/*
		LATER: many of these would be good for other text-based formats
		- trailing-newline: on/off/any/only
		- newline format: cr/crlf/lf/any (or dos/unix/mac?)
		- indent: tab/spaces/any
		- contains/doesnotcontain: specific text (license declaration/etc)
		- byte order mark: required/optional/forbidden
		- no homoglyphs
		- unicode: list of unicode character ranges allowed
			- predefined common sets
		- utf16
		- other charsets?
	*/
}

func textCheck(f *shared.FileContext) {

	raw, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}

	if textUtf8 {
		f.RecordResult("textUtf8", utf8.Valid(raw), nil)
	}

	if textAscii {
		asciiErr := IsAsciiBytes(raw)
		f.RecordResult("textAscii", asciiErr == nil, map[string]interface{}{
			"error": asciiErr,
		})
	}

	if !textNul {
		for index, b := range raw {
			if b == 0x00 {
				f.RecordResult("textNul", false, map[string]interface{}{
					"error":  shared.NewErrorAtPosition(fmt.Errorf("NUL byte"), index),
					"offset": index,
				})
				break
			}
		}
	}
}

func IsAsciiBytes(bytes []byte) error {
	for index, b := range bytes {
		if !IsAscii(b) {
			return shared.NewErrorAtPosition(fmt.Errorf("Non-ASCII value (0x%X)", b), index)
		}
	}
	return nil
}

func IsAscii(b byte) bool {
	if b > unicode.MaxASCII {
		return false
	}
	return true
}

func IsBadControl(b byte) bool {

	// not control
	if b >= ' ' {
		return false
	}

	// good control
	if b == '\t' || b == '\n' || b == '\r' {
		return false
	}

	return true
}
