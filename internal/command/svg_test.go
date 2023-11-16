package command

import (
	"testing"

	"github.com/FileFormatInfo/fflint/internal/shared"
)

func TestSvg(t *testing.T) {

	var fc = shared.FileContext{
		FilePath: "../../testdata/fflint128.svg",
	}
	//silent = true

	svgWidth.Set("64")
	svgCheck(&fc)
	if !fc.Success() {
		t.Errorf("width = 64")
	}
	fc.Reset()
	svgWidth.Set("64:")
	svgCheck(&fc)
	if !fc.Success() {
		t.Errorf("width >= 64")
	}
	fc.Reset()
	svgWidth.Set(":64")
	svgCheck(&fc)
	if !fc.Success() {
		t.Errorf("width <= 64")
	}

	// failing checks
	svgWidth.Set("16")
	svgCheck(&fc)
	if fc.Success() {
		t.Errorf("width = 16")
	}
	fc.Reset()
	svgWidth.Set("65:")
	svgCheck(&fc)
	if fc.Success() {
		t.Errorf("width >= 65")
	}
	fc.Reset()
	svgWidth.Set(":63")
	svgCheck(&fc)
	if fc.Success() {
		t.Errorf("width <= 63")
	}

	fc.Reset()
	svgWidth.Set("any")
	svgViewBox.Set("0,0,128,128")
	svgCheck(&fc)
	if !fc.Success() {
		t.Errorf("viewBox should work for 128")
	}

	fc.Reset()
	svgWidth.Set("any")
	svgViewBox.Set("0,0,64,64")
	svgCheck(&fc)
	if fc.Success() {
		t.Errorf("viewBox should not work for 64")
	}

	fc.FilePath = "../../testdata/fflint128.png"
	svgCheck(&fc)
	if fc.Success() {
		t.Errorf("invalid format")
	}
}
