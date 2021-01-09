# go-want

Go test library that delivers simple test assertions with useful verbose output.

## Usage

```go
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
	want.Eq(t, Abs(-1), 0) // will fail
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
	want.Eq(t, Title("hello\nfriendly\nworld"), "Hello\nfriendly\nWorld") // will fail
}
```

```
$ go test -v
--- FAIL: TestAbs (0.00s)
    readme_test.go:21: want.Eq(t, Abs(-1), 1): got 1
    readme_test.go:22: want.Eq(t, Abs(-1), 0): got 1, want 0
--- FAIL: TestTitle (0.00s)
    readme_test.go:41: want.Eq(t, Title("hello\nfriendly\nworld"), "Hello\nFriendly\nWorld"): got Hello
        Friendly
        World
    readme_test.go:42: want.Eq(t, Title("hello\nfriendly\nworld"), "Hello\nfriendly\nWorld"):
        --- Want
        +++ Got
        @@ -1,3 +1,3 @@
         Hello
        -friendly
        +Friendly
         World
```
