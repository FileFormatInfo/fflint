package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

// Range is an optional min/max pair
type Range struct {
	value    string
	hasStart bool
	start    uint64
	hasEnd   bool
	end      uint64
}

func (r *Range) String() string {
	if r.hasStart == false && r.hasEnd == false {
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

// Set will initialize the range
func (r *Range) Set(newValue string) error {
	r.value = newValue
	if len(newValue) == 0 || newValue == "any" {
		return nil
	}
	colonPos := strings.IndexByte(newValue, ':')
	if colonPos == -1 {
		v, err := strconv.ParseUint(newValue, 10, 64)
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
		v, err := strconv.ParseUint(newValue[1:], 10, 64)
		if err != nil {
			return err
		}
		r.hasStart = false
		r.hasEnd = true
		r.end = v
		return nil
	}
	if colonPos == len(newValue)-1 {
		v, err := strconv.ParseUint(newValue[:len(newValue)-1], 10, 64)
		if err != nil {
			return err
		}
		r.hasStart = true
		r.hasEnd = false
		r.start = v
		return nil
	}
	start, startErr := strconv.ParseUint(newValue[:colonPos], 10, 64)
	if startErr != nil {
		return startErr
	}
	end, endErr := strconv.ParseUint(newValue[colonPos+1:], 10, 64)
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
func (r *Range) Type() string {
	return "Range"
}

// Exists is true if there is a range to check
func (r *Range) Exists() bool {
	return r.value != "" && r.value != "any"
}

// Check if a value is within a range
func (r *Range) Check(v uint64) bool {
	if r.hasStart && v < r.start {
		return false
	}

	if r.hasEnd && v > r.end {
		return false
	}

	return true
}
