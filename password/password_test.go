package password

import (
	"testing"
)

func TestHashAndCheckPassword(t *testing.T) {
	cases := []string{
		"Testing",
		"Password123",
		"!ThErEHiTe$t",
		"Hello, 世界",
	}

	for _, c := range cases {
		got := HashPassword(c)
		if !CheckPassword(c, got) {
			t.Errorf("HashPassword&CheckPassword(%q) failed. Result: %q", c, got)
		}
	}
}
