package fuzzy

import (
	"fmt"
	"strings"

	"github.com/NectGmbH/match/fuzzy/dist"
)

const (
	defaultMinLength      = 2
	defaultMaxRelDistance = 0.2
)

// MatcherType denotes a fuzzy matcher instance
type MatcherType struct {
	distFn          dist.Fn
	minLength       int
	lowLengthAction bool
	maxRelDistance  float64

	isCaseSensitive bool
	isExact         bool
}

// Matcher instantiates a new Matcher
func Matcher() *MatcherType {
	return &MatcherType{
		minLength:      defaultMinLength,
		maxRelDistance: defaultMaxRelDistance,
		distFn:         dist.NewWagnerFischer().GenerateFn(),
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

// MaxRelativeDistance defines the maximum relative distance allowed for two elements to be considered a match
func (m *MatcherType) MaxRelativeDistance(maxDistance float64) *MatcherType {
	m.maxRelDistance = maxDistance

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
		return strings.EqualFold(reference, toMatch)
	}

	var distance float64
	if m.isCaseSensitive {
		distance = m.distFn(reference, toMatch)
	} else {
		distance = m.distFn(strings.ToLower(reference), strings.ToLower(toMatch))
	}

	if distance/float64(len(reference)) <= m.maxRelDistance {
		return true
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
		return false, fmt.Errorf("reference (%v) does not satisfy fmt.Stringer interface", reference)
	}

	// Assert b satisfies the fmt.Stringer interface
	bStringer, isStringer := toMatch.(fmt.Stringer)
	if !isStringer {
		return false, fmt.Errorf("rata to match (%v) does not satisfy fmt.Stringer interface", toMatch)
	}

	return m.MatchString(aStringer.String(), bStringer.String()), nil
}
