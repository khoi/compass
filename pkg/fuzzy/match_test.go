package fuzzy

import "testing"

func TestMatch(t *testing.T) {
	table := []struct {
		needle, haystack string
		expected         bool
	}{
		{"gosextant", "go/src/github.com/khoiln/sextant", true},
		{"Gosextant", "go/src/github.com/khoiln/sextant", false},
		{"srghub", "go/src/github.com/khoiln/sextant", true},
		{"hubgit", "go/src/github.com/khoiln/sextant", false},
		{"bò", "bún bò phở gà", true},
		{"bò gà", "bún bò phở gà", true},
		{"cơm", "bún bò phở gà", false},
		{"cơm bò", "cơm 👨‍👨‍👧‍👧 bò", true},
		{"cơgà", "cơm 👨‍👨‍👧‍👧 bò", false},
		{"👨bò", "cơm  👨‍👨‍👧‍👧 bò", true},
	}

	for _, c := range table {
		if output := Match(c.needle, c.haystack); output != c.expected {
			t.Errorf("Output: %v - Expected %v (for %s - %s)", output, c.expected, c.needle, c.haystack)
		}
	}
}

func TestMatchFold(t *testing.T) {
	table := []struct {
		needle, haystack string
		expected         bool
	}{
		{"gosextant", "go/src/github.com/khoiln/sextant", true},
		{"Gosextant", "go/src/github.com/khoiln/sextant", true},
		{"srgHub", "go/src/github.com/khoiln/sextant", true},
		{"hubgit", "go/src/github.com/khoiln/sextant", false},
		{"BÒ", "bún bò phở gà", true},
		{"bò Gà", "bún bò phở gà", true},
		{"cơM", "bún bò phở gà", false},
		{"cƠm BÒ", "cơm 👨‍👨‍👧‍👧 bò", true},
		{"cơGà", "cơm 👨‍👨‍👧‍👧 bò", false},
		{"👨bò", "cơm  👨‍👨‍👧‍👧 bò", true},
	}
	for _, c := range table {
		if output := MatchFold(c.needle, c.haystack); output != c.expected {
			t.Errorf("Output: %v - Expected %v (for %s - %s)", output, c.expected, c.needle, c.haystack)
		}
	}
}
