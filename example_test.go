package test_test

import (
	"testing"

	"4d63.com/test"
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
	test.Eq(t, Abs(-1), 1)
	// Output:
	// test.Eq(t, Abs(-1), 1): got 1
}

func ExampleEq_fail() {
	test.Eq(t, Abs(-1), 0)
	// Output:
	// test.Eq(t, Abs(-1), 0): got 1, want 0
}

func ExampleNotEq_pass() {
	test.NotEq(t, Abs(-1), -1)
	// Output:
	// test.NotEq(t, Abs(-1), -1): got 1, not -1
}

func ExampleNotEq_fail() {
	test.NotEq(t, Abs(-1), 1)
	// Output:
	// test.NotEq(t, Abs(-1), 1): got 1, want not 1
}

func ExampleNil_pass() {
	test.Nil(t, nil)
	// Output:
	// test.Nil(t, nil): got <nil>
}

func ExampleNil_fail() {
	test.Nil(t, thing)
	// Output:
	// test.Nil(t, thing): got 0, want <nil>
}

func ExampleNotNil_pass() {
	test.NotNil(t, thing)
	// Output:
	// test.NotNil(t, thing): got 0, not <nil>
}

func ExampleNotNil_fail() {
	test.NotNil(t, nil)
	// Output:
	// test.NotNil(t, nil): got <nil>, want not <nil>
}
