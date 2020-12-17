package fuzzy

import (
	"fmt"
	"strings"
)

const (
	defaultMinLength = 2

	defaultWeightInsert  = 1.
	defaultWeightReplace = 1.
	defaultWeightAdd     = 1.
)

// MatcherType denotes a fuzzy matcher instance
type MatcherType struct {
	weightInsert  float64
	weightReplace float64
	weightAdd     float64

	minLength           int
	lowLengthAction     bool
	maxWeightedDistance float64

	isCaseSensitive bool
	isExact         bool
}

// Matcher instantiates a new Matcher
func Matcher() *MatcherType {
	return &MatcherType{
		minLength:     defaultMinLength,
		weightInsert:  defaultWeightInsert,
		weightReplace: defaultWeightReplace,
		weightAdd:     defaultWeightAdd,
	}
}

// Exact defines the matcher to be exact, not fuzzy
func (m *MatcherType) Exact() *MatcherType {
	m.isExact = true

	return m
}

// CaseSensitive defines the matcher to be case sensitive
func (m *MatcherType) CaseSensitive() *MatcherType {
	m.isCaseSensitive = true

	return m
}

// MinLength defines the minimum reference / input string length to match successfully
func (m *MatcherType) MinLength(minLength int) *MatcherType {
	m.minLength = minLength

	return m
}

// LowLengthAction defines the match result in case input data is too short (as defined by minLength)
func (m *MatcherType) LowLengthAction(action bool) *MatcherType {
	m.lowLengthAction = action

	return m
}

// CaseSensitive defines the matcher to be case sensitive
func (m *MatcherType) MaxAbsoluteDistance(maxDistance float64) *MatcherType {
	m.isCaseSensitive = true

	return m
}

// Weights sets the weights / penalties for the Levenshtein distance calculation
func (m *MatcherType) Weights(weightInsert, weightReplace, weightAdd float64) *MatcherType {
	m.weightInsert = weightInsert
	m.weightReplace = weightReplace
	m.weightAdd = weightAdd

	return m
}

// MatchString performs a matching of two strings
func (m *MatcherType) MatchString(reference, toMatch string) bool {

	// In case either reference or string to match is too short,
	// return the defined result
	if len(reference) < m.minLength || len(toMatch) < m.minLength {
		return m.lowLengthAction
	}

	// Exact matching requested
	if m.isExact {

		// Case sensitive matching requested
		if m.isCaseSensitive {
			return reference == toMatch
		}

		// Case insensitive matching requested
		return strings.ToLower(reference) == strings.ToLower(toMatch)
	}

	return false
}

// MatchStringer performs a matching of two interfaces satisfying the fmt.Stringer interface
func (m *MatcherType) MatchStringer(reference, toMatch fmt.Stringer) bool {
	return m.MatchString(reference.String(), toMatch.String())
}

// Match performs a matching of two generic interfaces, if possible
func (m *MatcherType) Match(reference, toMatch interface{}) (bool, error) {

	// Assert a satisfies the fmt.Stringer interface
	aStringer, isStringer := reference.(fmt.Stringer)
	if !isStringer {
		return false, fmt.Errorf("Reference (%v) does not satisfy fmt.Stringer interface", reference)
	}

	// Assert b satisfies the fmt.Stringer interface
	bStringer, isStringer := toMatch.(fmt.Stringer)
	if !isStringer {
		return false, fmt.Errorf("Data to match (%v) does not satisfy fmt.Stringer interface", toMatch)
	}

	return m.MatchString(aStringer.String(), bStringer.String()), nil
}
