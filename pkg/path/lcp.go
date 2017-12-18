package path

import (
	"os"
	"path"
)

// LCP returns the longest common path for the list of path,
// or an empty string if no prefix is found.
func LCP(l []string) string {
	if len(l) == 0 {
		return ""
	}

	if len(l) == 1 {
		return path.Clean(l[0])
	}

	min := path.Clean(l[0])
	max := min

	for _, p := range l[1:] {
		p = path.Clean(p)

		switch {
		case p < min:
			min = p
		case p > max:
			max = p
		}
	}

	result := append([]byte(min), os.PathSeparator)
	for i := 0; i < len(result) && i < len(max); i++ {
		if result[i] != max[i] {
			result = result[:i]
			break
		}
	}

	for i := len(result) - 1; i >= 1; i-- {
		if result[i] == os.PathSeparator {
			result = result[:i]
			break
		}
	}

	return string(result)
}
