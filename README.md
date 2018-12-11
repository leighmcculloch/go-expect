# go-want

Go test library that delivers simple test assertions with useful verbose output.

## Usage

```go
import "4d63.com/want"

func TestAbs(t *testing.T) {
  want.Eq(t, Abs(-1), 1)
  want.Eq(t, Abs(-1), 0)
}
```

```
=== RUN   TestWantEq
--- FAIL: TestWantEq (0.00s)
    eq_test.go:11: got 1
    eq_test.go:12: got 1, want 0
```
