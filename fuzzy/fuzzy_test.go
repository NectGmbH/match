package fuzzy

import (
	"strings"
	"testing"

	"github.com/NectGmbH/match/normalize"
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
			t.Fatalf("Failed to match")
		}
	})
	t.Run("ICAO", func(t *testing.T) {
		if !Matcher().CaseSensitive().Exact().NormalizeFns(
			func(s string) string {
				return strings.ToUpper(s)
			},
			normalize.ReplaceUnicodeToICAO(),
		).MatchString("ABIJDE", "abĲde") {
			t.Fatalf("Failed to match")
		}
	})
}
