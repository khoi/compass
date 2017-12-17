package path

import (
	"testing"
)

var testTable = []struct {
	in  []string
	out string
}{
	{[]string{"foo"}, "foo"},
	{[]string{"/foo/bar/.."}, "/foo"},
	{[]string{"foo", "bar"}, ""},
	{[]string{"home/khoiln", "home/khoiln/foo", "home/khoiln/bar"}, "home/khoiln"},
	{[]string{"home/khoiln/bar/..", "home/khoiln/foo"}, "home/khoiln"},
	{[]string{"/abc/bcd/cdf", "/abc/bcd/cdf/foo", "/abc/bcd/chi/hij", "/abc/bcd/cdd"}, "/abc/bcd"},
	{[]string{"./abc/bcd/cdf", "./abc/bcd/cdf/foo", "./abc/bcd/chi/hij", "./abc/bcd/cdd"}, "abc/bcd"},
	{[]string{"/abc/bcd/cdf", "/"}, "/"},
}

func TestLCP(t *testing.T) {
	for _, c := range testTable {
		if out := LCP(c.in); out != c.out {
			t.Errorf("Expected %s - Got %s", c.out, out)
		}
	}
}
