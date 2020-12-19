package fuzzy

import (
	"testing"
)

func TestSimple(t *testing.T) {
	if !Matcher().CaseSensitive().Exact().MatchString("abc", "abc") {
		t.Fatalf("Failed to match simple identical string")
	}
}
