package path

import (
	"os"
	"path"
)

// LCP returns the longest common prefix for the list of path,
// or an empty string if no prefix is found.
func LCP(l []string) string {
	if len(l) == 0 {
		return ""
	}

	if len(l) == 1 {
		path.Clean(l[0])
	}

	c := []byte(path.Clean(l[0]))
	c = append(c, os.PathSeparator)

	for _, p := range l[1:] {
		p = path.Clean(p)

		if len(p) < len(c) {
			c = c[:len(p)]
		}

		for i := 0; i < len(c); i++ {
			if p[i] != c[i] {
				c = c[:i]
				break
			}
		}
	}

	for i := len(c) - 1; i >= 1; i-- {
		if c[i] == os.PathSeparator {
			c = c[:i]
			break
		}
	}

	return string(c)
}
