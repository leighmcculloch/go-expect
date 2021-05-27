package test_test

import (
	"testing"

	"4d63.com/test"
)

var (
	t     testing.TB = &printT{}
	thing            = 0
)

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

func ExampleEq_failDiff() {
	test.Eq(t, "Hello World\nG'day World\n", "Hello World\nG'day Mate")
	// Output:
	// test.Eq(t, "Hello World\nG'day World\n", "Hello World\nG'day Mate"):
	// --- Want
	// +++ Got
	// @@ -1,2 +1,3 @@
	//  Hello World
	// -G'day Mate
	// +G'day World
	// +
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

func ExampleEqJSON_pass() {
	test.EqJSON(
		t,
		[]byte(`{"key":"value1","key":"value2","key3":3}`),
		[]byte(` {
				"key":"value2",
				"key":"value1",
				"key3": 3
			}`),
	)
	// Output:
	// test.EqJSON(: got {
	//   "key": "value1",
	//   "key": "value2",
	//   "key3": 3
	// }
}

func ExampleEqJSON_fail() {
	test.EqJSON(
		t,
		[]byte(`{"key":"value1","key3":3}`),
		[]byte(` {
				"key3": 3
				"key":"value2",
			}`),
	)
	// Output:
	// test.EqJSON(:
	// --- Want
	// +++ Got
	// @@ -1,5 +1,5 @@
	//  {
	// -  "key": "value2",
	// +  "key": "value1",
	//    "key3": 3
	//  }
}
