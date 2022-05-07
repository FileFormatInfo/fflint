package shared

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/mitchellh/go-homedir"
)

var (
	bar *pb.ProgressBar
)

var barIntFormat string
var cwd string
var userdir string

// ProgressCount displays a counter while calculating the real total progress
var (
	fileCount int
)

// ProgressCount called when loading the list (before end is known)
func ProgressCount() {
	if progress {
		// display count of scanned file names
		fileCount++
		saveCursorPosition()
		fmt.Fprintf(os.Stderr, "Scanning (%d)...", fileCount)
		restoreCursorPosition()
	}
}

// ProgressStart start the progress bar
func ProgressStart(fcs []FileContext) {
	if progress {
		var max int64
		if true {
			for _, fc := range fcs {
				fi, _ := fc.Stat()
				max += fi.Size()
			}
		} else {
			max = int64(len(fcs))
		}
		pb.RegisterElement("mypercent", myPercentElement, false)

		numWidth := len(fmt.Sprintf("%d", len(fcs)))
		barIntFormat = fmt.Sprintf("%%%dd", numWidth)
		path, err := os.Getwd()
		if err != nil {
			if Debug {
				fmt.Fprintf(os.Stderr, "ERROR: unable to retrieve current working directory (%s)\n", err.Error())
			}
		} else {
			cwd = path
		}
		userdir, _ = homedir.Dir()

		pb.RegisterElement("fixedbar", pb.ElementBar, false)
		tmpl := `PASS {{string . "good"}} / FAIL {{string . "bad"}} {{ fixedbar . "[" "=" ">" "·" "]"}} {{mypercent . }} {{string . "filename"}}`
		bar = pb.Start64(max)
		bar.SetTemplateString(tmpl)
	}
}

// ProgressUpdate update the progress bar
func ProgressUpdate(good int, bad int, fc FileContext) {
	if bar != nil {
		bar.Set("good", fmt.Sprintf(barIntFormat, good))
		bar.Set("bad", fmt.Sprintf(barIntFormat, bad))
		if fc.FilePath == "" {
			elapsed := time.Now().Truncate(time.Second).Sub(bar.StartTime().Truncate(time.Second))
			if elapsed < 10*time.Second {
				elapsed = time.Now().Truncate(time.Millisecond).Sub(bar.StartTime().Truncate(time.Millisecond))
			}
			bar.Set("filename", fmt.Sprintf("Done in %s", elapsed))
		} else {
			if true {
				fi, _ := fc.Stat()
				bar.Add64(fi.Size())
			} else {
				bar.Add64(1)
			}
			if strings.HasPrefix(fc.FilePath, cwd) {
				bar.Set("filename", fmt.Sprintf(".%s", fc.FilePath[len(cwd):]))
			} else if strings.HasPrefix(fc.FilePath, userdir) {
				bar.Set("filename", fmt.Sprintf("~%s", fc.FilePath[len(userdir):]))
			} else {
				bar.Set("filename", fc.FilePath)
			}
		}
	}
}

// ProgressEnd call when completely done
func ProgressEnd() {
	if bar != nil {
		bar.Finish()
		if Debug {
			fmt.Fprintf(os.Stderr, "DEBUG: Bytes processed: %d\n", bar.Total())
		}
	}
}

// variation of pb.ElementPercent that is fixed width and rounds down
var myPercentElement pb.ElementFunc = func(state *pb.State, args ...string) string {
	if state.Total() == 0 {
		return "      "
	}

	percent1000 := state.Value() * 1000 / state.Total()

	return fmt.Sprintf("%3d.%01d%%", percent1000/10, percent1000%10)
}

func clearLine() {
	fmt.Fprintf(os.Stderr, "\033[2K")
}
func saveCursorPosition() {
	fmt.Fprintf(os.Stderr, "\033[s")
}

func restoreCursorPosition() {
	fmt.Fprintf(os.Stderr, "\033[u")
}
