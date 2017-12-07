package fuzzy

import "unicode"

func Match(needle, haystack string) bool {
	return match(needle, haystack, func(r1 rune, r2 rune) bool {
		return r1 == r2
	})
}

func MatchFold(needle, haystack string) bool {
	return match(needle, haystack, func(r1 rune, r2 rune) bool {
		return unicode.ToLower(r1) == unicode.ToLower(r2)
	})
}

func match(needle, haystack string, compare func(rune, rune) bool) bool {
	if len(needle) > len(haystack) {
		return false
	}
	if len(needle) == len(haystack) {
		return needle == haystack
	}

	n, h := []rune(needle), []rune(haystack)

	nIdx, hIdx := 0, 0
	for nIdx < len(n) {
		if hIdx >= len(h) {
			return false
		}
		if compare(n[nIdx], h[hIdx]) {
			nIdx++
		}
		hIdx++
	}
	return true
}
