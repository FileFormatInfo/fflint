package cmd

import (
	"github.com/cheggaaa/pb/v3"
)

var (
	bar *pb.ProgressBar
)

func ProgressStart(max int) {
	if progress {
		bar = pb.StartNew(max)
	}
}

func ProgressUpdate(sucess bool) {
	if bar != nil {
		bar.Increment()
	}
}

func ProgressEnd() {
	if bar != nil {
		bar.Finish()
	}
}
