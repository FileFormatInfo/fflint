package cmd

import (
	"strings"
	"testing"

	"github.com/shopspring/decimal"
)

func singleDecimalRangeArrayTest(t *testing.T, rs string, vs string, desired bool) {
	ra := DecimalRangeArray{}

	va := strings.Split(vs, ",")
	da := make([]decimal.Decimal, len(va))
	for i, v := range va {
		da[i] = decimal.RequireFromString(v)
	}

	ra.Set(rs)

	actual := ra.Check(da)

	if actual != desired {
		t.Errorf("%s for %s: actual=%t, expected=%t", rs, vs, actual, desired)
	}
}

func TestDecimalRangeArray(t *testing.T) {
	singleDecimalRangeArrayTest(t, "0,0,120,60", "0,0,120,60", true)
	singleDecimalRangeArrayTest(t, "0,0,120,60", "0,0,121,60", false)
	singleDecimalRangeArrayTest(t, "0,0,120,60", "0,0,120,61", false)
	singleDecimalRangeArrayTest(t, "0,0,120,60", "0,0,121,61", false)

	singleDecimalRangeArrayTest(t, "0,0,119:121,60", "0,0,120,60", true)
	singleDecimalRangeArrayTest(t, "0,0,119:121,60", "0,0,121,60", true)
	singleDecimalRangeArrayTest(t, "0,0,119:121,60", "0,0,119,60", true)
	singleDecimalRangeArrayTest(t, "0,0,119:121,60", "0,0,122,60", false)
	singleDecimalRangeArrayTest(t, "0,0,119:121,60", "0,0,118,60", false)

	singleDecimalRangeArrayTest(t, "0,0,119.5:120.5,60", "0,0,120,60", true)
	singleDecimalRangeArrayTest(t, "0,0,119.5:120.5,60", "0,0,120.5,60", true)
	singleDecimalRangeArrayTest(t, "0,0,119.5:120.5,60", "0,0,119.5,60", true)
	singleDecimalRangeArrayTest(t, "0,0,119.5:120.5,60", "0,0,120.6,60", false)
	singleDecimalRangeArrayTest(t, "0,0,119.5:120.5,60", "0,0,119.4,60", false)

}
