package cmd

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// DecimalRange is an optional min/max pair
type DecimalRange struct {
	value    string
	hasStart bool
	start    decimal.Decimal
	hasEnd   bool
	end      decimal.Decimal
}

func (r *DecimalRange) String() string {
	if r.hasStart == false && r.hasEnd == false {
		return "any"
	}
	if r.hasStart == false {
		return fmt.Sprintf("<=%s", r.end)
	}
	if r.hasEnd == false {
		return fmt.Sprintf(">=%s", r.start)
	}
	if r.start == r.end {
		return fmt.Sprintf("=%s", r.start)
	}

	return r.value
}

// Set will initialize the range
func (r *DecimalRange) Set(newValue string) error {
	r.value = newValue
	if len(newValue) == 0 || newValue == "any" {
		return nil
	}
	colonPos := strings.IndexByte(newValue, ':')
	if colonPos == -1 {
		v, err := decimal.NewFromString(newValue)
		if err != nil {
			return err
		}
		r.hasStart = true
		r.hasEnd = true
		r.start = v
		r.end = v
		return nil
	}
	if colonPos == 0 {
		v, err := decimal.NewFromString(newValue[1:])
		if err != nil {
			return err
		}
		r.hasStart = false
		r.hasEnd = true
		r.end = v
		return nil
	}
	if colonPos == len(newValue)-1 {
		v, err := decimal.NewFromString(newValue[:len(newValue)-1])
		if err != nil {
			return err
		}
		r.hasStart = true
		r.hasEnd = false
		r.start = v
		return nil
	}
	start, startErr := decimal.NewFromString(newValue[:colonPos])
	if startErr != nil {
		return startErr
	}
	end, endErr := decimal.NewFromString(newValue[colonPos+1:])
	if endErr != nil {
		return endErr
	}
	r.hasStart = true
	r.hasEnd = true
	r.start = start
	r.end = end
	return nil
}

// Type is a description of range
func (r *DecimalRange) Type() string {
	return "DecimalRange"
}

// Exists is true if there is a range to check
func (r *DecimalRange) Exists() bool {
	return r.value != "" && r.value != "any"
}

// Check if a value is within a range
func (r *DecimalRange) Check(v decimal.Decimal) bool {
	if r.hasStart && v.LessThan(r.start) {
		return false
	}

	if r.hasEnd && v.GreaterThan(r.end) {
		return false
	}

	return true
}

type ViewboxRange struct {
	ranges [4]DecimalRange
}

// LATER: String, Set, Exists, Check
