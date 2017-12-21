package path

import (
	"testing"
)

var testTable = []struct {
	in  []string
	out string
}{
	{[]string{}, ""},
	{[]string{"foo"}, "foo"},
	{[]string{"/foo/bar/.."}, "/foo"},
	{[]string{"foo", "bar"}, ""},
	{[]string{"home/khoiracle", "home/khoiracle/foo", "home/khoiracle/bar"}, "home/khoiracle"},
	{[]string{"home/khoiracle/bar/..", "home/khoiracle/foo"}, "home/khoiracle"},
	{[]string{"/abc/bcd/cdf", "/abc/bcd/cdf/foo", "/abc/bcd/chi/hij", "/abc/bcd/cdd"}, "/abc/bcd"},
	{[]string{"./abc/bcd/cdf", "./abc/bcd/cdf/foo", "./abc/bcd/chi/hij", "./abc/bcd/cdd"}, "abc/bcd"},
	{[]string{"/abc/bcd/cdf", "/"}, "/"},
	{[]string{"/abc/def/ghj", "/abc/def"}, "/abc/def"},
	{[]string{"Github/khoi/ios", "Github/khoi/webcontent-ios", "Github/khoi/ios/iosNetworking"}, "Github/khoi"},
}

func TestLCP(t *testing.T) {
	for _, c := range testTable {
		if out := LCP(c.in); out != c.out {
			t.Errorf("Expected %s - Got %s", c.out, out)
		}
	}
}
