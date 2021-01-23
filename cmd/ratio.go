package cmd

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

// Ratio is an optional min/max pair
type Ratio struct {
	value    string
	hasStart bool
	start    decimal.Decimal
	hasEnd   bool
	end      decimal.Decimal
	hasRat   bool
	exact    big.Rat
}

func (r *Ratio) String() string {
	if r.hasRat == false && r.hasStart == false && r.hasEnd == false {
		return "any"
	}
	if r.hasStart == false {
		return fmt.Sprintf("<=%d", r.end)
	}
	if r.hasEnd == false {
		return fmt.Sprintf(">=%d", r.start)
	}
	if r.start == r.end {
		return fmt.Sprintf("=%d", r.start)
	}

	return r.value
}

// Set will initialize the ratio
func (r *Ratio) Set(newValue string) error {
	r.value = newValue
	if len(newValue) == 0 || newValue == "any" {
		return nil
	}
	colonPos := strings.IndexByte(newValue, ':')
	if colonPos == -1 {
		slashPos := strings.IndexByte(newValue, '/')
		if slashPos == -1 {
			v, err := decimal.NewFromString(newValue)
			if err != nil {
				return err
			}
			r.hasStart = true
			r.hasEnd = true
			r.start = v
			r.end = v
		} else {
			r.hasRat = true
			//r.exact = big.NewRat(x, y)
		}
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

// Type is a description of ratio
func (r *Ratio) Type() string {
	return "Ratio"
}

// Exists is true if there is a ratio to check
func (r *Ratio) Exists() bool {
	return r.value != "" && r.value != "any"
}

// Check if a value is within a ratio
func (r *Ratio) Check(v decimal.Decimal) bool {
	if r.hasStart && v.LessThan(r.start) {
		return false
	}

	if r.hasEnd && v.GreaterThan(r.end) {
		return false
	}

	return true
}
