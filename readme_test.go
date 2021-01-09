// +build readme_example

package want_test

import (
	"strings"
	"testing"
	"unicode"

	"4d63.com/want"
)

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func TestAbs(t *testing.T) {
	want.Eq(t, Abs(-1), 1)
	want.Eq(t, Abs(-1), 0)
}

func Title(s string) string {
	sb := strings.Builder{}
	beginningWord := true
	for _, r := range s {
		if beginningWord {
			r = unicode.ToUpper(r)
			beginningWord = false
		} else if unicode.IsSpace(r) {
			beginningWord = true
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

func TestTitle(t *testing.T) {
	want.Eq(t, Title("hello\nfriendly\nworld"), "Hello\nFriendly\nWorld")
	want.Eq(t, Title("hello\nfriendly\nworld"), "Hello\nfriendly\nWorld")
}
