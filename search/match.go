package search

func Find(needle, haystack string) bool {
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
		if n[nIdx] == h[(hIdx)] {
			nIdx++
		}
		hIdx++
	}
	return true
}
