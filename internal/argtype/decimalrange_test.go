package argtype

import (
	"testing"

	"github.com/shopspring/decimal"
)

func singleDecimalRangeTest(t *testing.T, rs string, vs string, desired bool) {
	r := DecimalRange{}

	v := decimal.RequireFromString(vs)

	r.Set(rs)

	actual := r.Check(v)

	if actual != desired {
		t.Errorf("%s for %s: actual=%t, expected=%t", rs, v, actual, desired)
	}
}

func TestDecimalRange(t *testing.T) {
	singleDecimalRangeTest(t, "100", "100", true)
	singleDecimalRangeTest(t, "100", "101", false)
	singleDecimalRangeTest(t, "100", "99", false)

	singleDecimalRangeTest(t, "100-", "100", true)
	singleDecimalRangeTest(t, "100-", "101", true)
	singleDecimalRangeTest(t, "100-", "99", false)

	singleDecimalRangeTest(t, "-100", "100", true)
	singleDecimalRangeTest(t, "-100", "101", false)
	singleDecimalRangeTest(t, "-100", "99", true)

	singleDecimalRangeTest(t, "100-100", "100", true)
	singleDecimalRangeTest(t, "100-100", "101", false)
	singleDecimalRangeTest(t, "100-100", "99", false)

	singleDecimalRangeTest(t, "99-101", "100", true)
	singleDecimalRangeTest(t, "99-101", "101", true)
	singleDecimalRangeTest(t, "99-101", "99", true)
	singleDecimalRangeTest(t, "99-101", "102", false)
	singleDecimalRangeTest(t, "99-101", "98", false)
	//
	singleDecimalRangeTest(t, "100", "100.0", true)
	singleDecimalRangeTest(t, "100", "101.0", false)
	singleDecimalRangeTest(t, "100", "99.0", false)

	singleDecimalRangeTest(t, "100-", "100.0", true)
	singleDecimalRangeTest(t, "100-", "101.0", true)
	singleDecimalRangeTest(t, "100-", "99.0", false)

	singleDecimalRangeTest(t, "-100", "100.0", true)
	singleDecimalRangeTest(t, "-100", "101.0", false)
	singleDecimalRangeTest(t, "-100", "99.0", true)

	singleDecimalRangeTest(t, "100-100", "100.0", true)
	singleDecimalRangeTest(t, "100-100", "101.0", false)
	singleDecimalRangeTest(t, "100-100", "99.0", false)

	singleDecimalRangeTest(t, "99-101", "100.0", true)
	singleDecimalRangeTest(t, "99-101", "101.0", true)
	singleDecimalRangeTest(t, "99-101", "100.9", true)
	singleDecimalRangeTest(t, "99-101", "101.1", false)
	singleDecimalRangeTest(t, "99-101", "99.0", true)
	singleDecimalRangeTest(t, "99-101", "102.0", false)
	singleDecimalRangeTest(t, "99-101", "98.0", false)
	//
	singleDecimalRangeTest(t, "100.0", "100.0", true)
	singleDecimalRangeTest(t, "100.0", "101.0", false)
	singleDecimalRangeTest(t, "100.0", "99.0", false)

	singleDecimalRangeTest(t, "100.0-", "100.0", true)
	singleDecimalRangeTest(t, "100.0-", "101.0", true)
	singleDecimalRangeTest(t, "100.0-", "99.0", false)

	singleDecimalRangeTest(t, "-100.0", "100.0", true)
	singleDecimalRangeTest(t, "-100.0", "101.0", false)
	singleDecimalRangeTest(t, "-100.0", "99.0", true)

	singleDecimalRangeTest(t, "100.0-100.0", "100.0", true)
	singleDecimalRangeTest(t, "100.0-100.0", "101.0", false)
	singleDecimalRangeTest(t, "100.0-100.0", "99.0", false)

	singleDecimalRangeTest(t, "99.0-101.0", "100.0", true)
	singleDecimalRangeTest(t, "99.0-101.0", "101.0", true)
	singleDecimalRangeTest(t, "99.0-101.0", "100.9", true)
	singleDecimalRangeTest(t, "99.0-101.0", "101.1", false)
	singleDecimalRangeTest(t, "99.0-101.0", "99.0", true)
	singleDecimalRangeTest(t, "99.0-101.0", "102.0", false)
	singleDecimalRangeTest(t, "99.0-101.0", "98.0", false)

}
