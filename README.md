# 4d63.com/test

[![Go Reference](https://pkg.go.dev/badge/4d63.com/test.svg)](https://pkg.go.dev/4d63.com/test)

Go test library that delivers simple test assertions with useful verbose output.

## Features

 - [x] Lightweight
 - [x] Simplicity of [testify](https://github.com/stretchr/testify)
 - [x] Verbose logs based on successful assertions
 - [x] Concise git diff comparisons
 - [x] JSON comparisons (powered by [pretty](https://github.com/tidwall/pretty))

## Todo
 - [ ] Error is/as comparisons
 - [ ] Slice contains
 - [ ] Map contains
 - [ ] Regex comparisons

## Usage

```go
import (
	"strings"
	"testing"
	"unicode"

	"4d63.com/test"
)

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func TestAbs(t *testing.T) {
	test.Eq(t, Abs(-1), 1)
	test.Eq(t, Abs(-1), 0) // will fail
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
	test.Eq(t, Title("hello\nfriendly\nworld"), "Hello\nFriendly\nWorld")
	test.Eq(t, Title("hello\nfriendly\nworld"), "Hello\nfriendly\nWorld") // will fail
}
```

```
$ go test -v
--- FAIL: TestAbs (0.00s)
    readme_test.go:21: test.Eq(t, Abs(-1), 1): got 1
    readme_test.go:22: test.Eq(t, Abs(-1), 0): got 1, want 0
--- FAIL: TestTitle (0.00s)
    readme_test.go:41: test.Eq(t, Title("hello\nfriendly\nworld"), "Hello\nFriendly\nWorld"): got Hello
        Friendly
        World
    readme_test.go:42: test.Eq(t, Title("hello\nfriendly\nworld"), "Hello\nfriendly\nWorld"):
        --- Want
        +++ Got
        @@ -1,3 +1,3 @@
         Hello
        -friendly
        +Friendly
         World
```
