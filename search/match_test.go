package search

import "testing"

func TestFind(t *testing.T) {
	table := []struct {
		needle, haystack string
		expected         bool
	}{
		{"gosextant", "go/src/github.com/khoiln/sextant", true},
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
		if output := Find(c.needle, c.haystack); output != c.expected {
			t.Errorf("Output: %v - Expected %v (for %s - %s)", output, c.expected, c.needle, c.haystack)
		}
	}
}
