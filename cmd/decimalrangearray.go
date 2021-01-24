package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/shopspring/decimal"
)

// DecimalRangeArray checks a set of DecimalRanges
type DecimalRangeArray struct {
	ranges []DecimalRange
}

func (ra *DecimalRangeArray) String() string {
	var sb strings.Builder

	for i, r := range ra.ranges {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(r.String())
	}
	return sb.String()
}

// Type is a description of range
func (ra *DecimalRangeArray) Type() string {
	return "DecimalRangeArray"
}

// Exists is true if there is a range to check
func (ra *DecimalRangeArray) Exists() bool {
	for _, r := range ra.ranges {
		if r.Exists() {
			return true
		}
	}
	return false
}

// Check if a value is within a range
func (ra *DecimalRangeArray) Check(va []decimal.Decimal) bool {

	if len(va) != len(ra.ranges) {
		return false
	}

	for i, v := range va {
		if !ra.ranges[i].Check(v) {
			return false
		}
	}

	return true
}

// CheckString if a value is within a range
func (ra *DecimalRangeArray) CheckString(s, sep string) bool {

	va := strings.Split(s, sep)
	da := make([]decimal.Decimal, len(va))
	for i, v := range va {
		d, err := decimal.NewFromString(v)
		if err != nil {
			if debug {
				fmt.Fprintf(os.Stderr, "Unable to parse %s as a decimal (#%d in %s)\n", v, i, s)
			}
			return false
		}
		da[i] = d
	}

	return ra.Check(da)
}

// Set sets the value of a DecimalRangeArray
func (ra *DecimalRangeArray) Set(newValue string) error {
	va := strings.Split(newValue, ",")

	ra.ranges = make([]DecimalRange, len(va))

	for i, v := range va {
		r := DecimalRange{}
		err := r.Set(v)
		if err != nil {
			return err
		}
		ra.ranges[i] = r
	}
	return nil
}
