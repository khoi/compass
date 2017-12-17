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

	min := []byte(path.Clean(l[0]))
	min = append(min, os.PathSeparator)
	max := min

	for _, p := range l[1:] {
		p = path.Clean(p)

		switch {
		case len(p) < len(min):
			min = []byte(p)
		case len(p) > len(min):
			max = []byte(p)
		}
	}

	for i := 0; i < len(min) && i < len(max); i++ {
		if min[i] != max[i] {
			min = min[:i]
			break
		}
	}

	for i := len(min) - 1; i >= 1; i-- {
		if min[i] == os.PathSeparator {
			min = min[:i]
			break
		}
	}

	return string(min)
}
