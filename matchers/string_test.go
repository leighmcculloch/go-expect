package matchers

import (
	"fmt"
	"testing"
)

func TestStringEq(t *testing.T) {
	cases := []struct {
		e string
		a string
		r bool
	}{
		{"", "", true},
		{"1", "1", true},
		{"1hello", "1hello", true},
		{"hello", "1hello", false},
		{"1hello", "hello", false},
		{"1helo", "1hello", false},
		{"1hello", "1helo", false},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%s, %s, %t", c.a, c.e, c.r), func(t *testing.T) {
			r := c.a == c.e
			if r != c.r {
				t.Errorf("%q == %q got %t, want %t", c.a, c.e, c.r, r)
			}
		})
	}
}
