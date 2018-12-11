package want_test

import (
	"testing"

	"4d63.com/want"
)

var t testing.TB = &fakeT{}
var thing = 0

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func ExampleEq() {
	want.Eq(t, Abs(-1), 1)
	// Output:
}

func ExampleNotEq() {
	want.NotEq(t, Abs(-1), -1)
	// Output:
}

func ExampleNil() {
	want.Nil(t, thing)
	// Output:
}

func ExampleNotNil() {
	want.NotNil(t, thing)
	// Output:
}
