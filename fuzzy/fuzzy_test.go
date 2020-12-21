package fuzzy

import (
	"testing"
)

func TestSimple(t *testing.T) {
	t.Run("Exact", func(t *testing.T) {
		if !Matcher().CaseSensitive().Exact().MatchString("abc", "abc") {
			t.Fatalf("Failed to match")
		}
	})
	t.Run("CaseInsensitive", func(t *testing.T) {
		if !Matcher().MatchString("abcde", "ABCD") {
			t.Fatalf("Failed to match")
		}
	})
	t.Run("CaseSensitive", func(t *testing.T) {
		if !Matcher().CaseSensitive().MatchString("abcde", "abcd") {
			t.Fatalf("Failed to matchg")
		}
	})

}
