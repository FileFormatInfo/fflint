package cmd

import (
	"testing"
)

func singleRangeTest(t *testing.T, rs string, v uint64, desired bool) {
	r := Range{}

	r.Set(rs)

	actual := r.Check(v)

	if actual != desired {
		t.Errorf("%s for %d: actual=%t, expected=%t", rs, v, actual, desired)
	}
}

func TestRange(t *testing.T) {
	singleRangeTest(t, "100", 100, true)
	singleRangeTest(t, "100", 101, false)
	singleRangeTest(t, "100", 99, false)

	singleRangeTest(t, "100:", 100, true)
	singleRangeTest(t, "100:", 101, true)
	singleRangeTest(t, "100:", 99, false)

	singleRangeTest(t, ":100", 100, true)
	singleRangeTest(t, ":100", 101, false)
	singleRangeTest(t, ":100", 99, true)

	singleRangeTest(t, "100:100", 100, true)
	singleRangeTest(t, "100:100", 101, false)
	singleRangeTest(t, "100:100", 99, false)

	singleRangeTest(t, "99:101", 100, true)
	singleRangeTest(t, "99:101", 101, true)
	singleRangeTest(t, "99:101", 99, true)
	singleRangeTest(t, "99:101", 102, false)
	singleRangeTest(t, "99:101", 98, false)
}
