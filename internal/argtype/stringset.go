package argtype

import (
	"fmt"
	"sort"

	"golang.org/x/exp/maps"
)

// string flag with value in a set of acceptable strings
type StringSet struct {
	name         string
	defaultValue string
	allowed      map[string]bool
	currentValue string
}

func NewStringSet(name, defaultValue string, values []string) StringSet {
	allowed := make(map[string]bool)
	for _, value := range values {
		allowed[value] = true
	}
	return StringSet{
		name:         name,
		defaultValue: defaultValue,
		allowed:      allowed,
		currentValue: "",
	}
}

func (ss *StringSet) String() string {
	if ss.currentValue != "" {
		return ss.currentValue
	}
	return ss.defaultValue
}

func (ss *StringSet) HelpText() string {

	keys := maps.Keys(ss.allowed)
	sort.Strings(keys)
	retVal := "[ "
	for i, k := range keys {
		if i > 0 {
			retVal += " &#x7c; "
		}
		retVal += k
	}
	retVal += " ]"

	return retVal
}

// Set will initialize the StringSet
func (ss *StringSet) Set(newValue string) error {
	_, exists := ss.allowed[newValue]
	if !exists {
		return fmt.Errorf("Invalid setting '%s' for %s", newValue, ss.name)
	}
	ss.currentValue = newValue
	return nil
}

// Type is the description of this StringSet
func (ss *StringSet) Type() string {
	return ss.name
}

// Exists is true if there is a range to check
func (ss *StringSet) Exists() bool {
	return ss.currentValue != ""
}
