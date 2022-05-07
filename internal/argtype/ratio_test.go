package argtype

import (
	"testing"

	"github.com/shopspring/decimal"
)

func singleRatioTest(t *testing.T, rs string, v decimal.Decimal, desired bool) {
	r := Ratio{}

	r.Set(rs)

	actual := r.Check(v)

	if actual != desired {
		t.Errorf("setting=%s value=%s: actual=%t, expected=%t", rs, v, actual, desired)
	}
}

// LATER: move decimal stuff into alternate ratio.CheckInt
func singleRatioTestInt(t *testing.T, rs string, x int64, y int64, desired bool) {
	v := decimal.NewFromInt(x).Div(decimal.NewFromInt(y))

	singleRatioTest(t, rs, v, desired)
}

func singleRatioTestFloat(t *testing.T, rs string, x float64, y float64, desired bool) {
	v := decimal.NewFromFloat(x).Div(decimal.NewFromFloat(y))

	singleRatioTest(t, rs, v, desired)
}

func TestRatio(t *testing.T) {
	singleRatioTestInt(t, "1", 1, 1, true)
	singleRatioTestInt(t, "1", 1, 2, false)
	singleRatioTestInt(t, "1", 2, 1, false)
	singleRatioTestInt(t, "1", 2, 1, false)
	singleRatioTestInt(t, "1", 2, 2, true)

	singleRatioTestInt(t, "1.0", 1, 1, true)
	singleRatioTestInt(t, "1.0", 1, 2, false)
	singleRatioTestInt(t, "1.0", 2, 1, false)
	singleRatioTestInt(t, "1.0", 2, 1, false)
	singleRatioTestInt(t, "1.0", 2, 2, true)

	singleRatioTestFloat(t, "1.0", 1.0, 1.0, true)
	singleRatioTestFloat(t, "1.0", 1.0, 2.0, false)
	singleRatioTestFloat(t, "1.0", 2.0, 1.0, false)
	singleRatioTestFloat(t, "1.0", 2.0, 1.0, false)
	singleRatioTestFloat(t, "1.0", 2.0, 2.0, true)

	singleRatioTestFloat(t, "1", 1.0, 1.0, true)
	singleRatioTestFloat(t, "1", 1.0, 2.0, false)
	singleRatioTestFloat(t, "1", 2.0, 1.0, false)
	singleRatioTestFloat(t, "1", 2.0, 1.0, false)
	singleRatioTestFloat(t, "1", 2.0, 2.0, true)

	//LATER: support for exact specified with ratio like "4/3"
	//LATER: support for a list of values like "16/9,4/3,1"
}
