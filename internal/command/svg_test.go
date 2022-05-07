package command

import (
	"testing"
)

func TestSvg(t *testing.T) {

	var fc = FileContext{
		FilePath: "../tests/badger128.svg",
	}
	//silent = true

	svgWidth.Set("64")
	svgCheck(&fc)
	if !fc.success() {
		t.Errorf("width = 64")
	}
	fc.reset()
	svgWidth.Set("64:")
	svgCheck(&fc)
	if !fc.success() {
		t.Errorf("width >= 64")
	}
	fc.reset()
	svgWidth.Set(":64")
	svgCheck(&fc)
	if !fc.success() {
		t.Errorf("width <= 64")
	}

	// failing checks
	svgWidth.Set("16")
	svgCheck(&fc)
	if fc.success() {
		t.Errorf("width = 16")
	}
	fc.reset()
	svgWidth.Set("65:")
	svgCheck(&fc)
	if fc.success() {
		t.Errorf("width >= 65")
	}
	fc.reset()
	svgWidth.Set(":63")
	svgCheck(&fc)
	if fc.success() {
		t.Errorf("width <= 63")
	}

	fc.reset()
	svgWidth.Set("any")
	svgViewBox.Set("0,0,128,128")
	svgCheck(&fc)
	if !fc.success() {
		t.Errorf("viewBox should work for 128")
	}

	fc.reset()
	svgWidth.Set("any")
	svgViewBox.Set("0,0,64,64")
	svgCheck(&fc)
	if fc.success() {
		t.Errorf("viewBox should not work for 64")
	}

	fc.FilePath = "../tests/badger128.png"
	svgCheck(&fc)
	if fc.success() {
		t.Errorf("invalid format")
	}
}
