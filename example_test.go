// +build !readme_example

package want_test

import (
	"testing"

	"4d63.com/want"
)

var t testing.TB = &printT{}
var thing = 0

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func ExampleEq_pass() {
	want.Eq(t, Abs(-1), 1)
	// Output:
	// want.Eq(t, Abs(-1), 1): got 1
}

func ExampleEq_fail() {
	want.Eq(t, Abs(-1), 0)
	// Output:
	// want.Eq(t, Abs(-1), 0): got 1, want 0
}

func ExampleNotEq_pass() {
	want.NotEq(t, Abs(-1), -1)
	// Output:
	// want.NotEq(t, Abs(-1), -1): got 1, not -1
}

func ExampleNotEq_fail() {
	want.NotEq(t, Abs(-1), 1)
	// Output:
	// want.NotEq(t, Abs(-1), 1): got 1, want not 1
}

func ExampleNil_pass() {
	want.Nil(t, nil)
	// Output:
	// want.Nil(t, nil): got <nil>
}

func ExampleNil_fail() {
	want.Nil(t, thing)
	// Output:
	// want.Nil(t, thing): got 0, want <nil>
}

func ExampleNotNil_pass() {
	want.NotNil(t, thing)
	// Output:
	// want.NotNil(t, thing): got 0, not <nil>
}

func ExampleNotNil_fail() {
	want.NotNil(t, nil)
	// Output:
	// want.NotNil(t, nil): got <nil>, want not <nil>
}
