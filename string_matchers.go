package expect

import (
	"4d63.com/expect/matchers"
)

func (s String) ToEq(expected string) bool {
	s.x.tb.Helper()
	return s.To(matchers.StringEq(expected), "got %q, want %q", s.s, expected)
}

func (s String) NotToEq(expected string) bool {
	s.x.tb.Helper()
	return s.NotTo(matchers.StringEq(expected), "got %q, did not want %q", s.s, expected)
}
