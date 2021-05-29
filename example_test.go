package test_test

import (
	"testing"

	"4d63.com/test"
)

var t testing.TB = &printT{}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func ExampleEqual_pass() {
	test.Equal(t, Abs(-1), 1)
	// Output:
	// test.Equal(t, Abs(-1), 1): got 1
}

func ExampleEqual_fail() {
	test.Equal(t, Abs(-1), 0)
	// Output:
	// test.Equal(t, Abs(-1), 0): got 1, want 0
}

func ExampleEqual_failComparingDifferentTypes() {
	test.Equal(t, 1, 1.0)
	// Output:
	// test.Equal(t, 1, 1.0):
	// --- got
	// +++ want
	// @@ -1,2 +1,2 @@
	// -(int) 1
	// +(float64) 1
}

func ExampleEqual_failDiff() {
	test.Equal(t, "Hello World\nG'day World\n", "Hello World\nG'day Mate")
	// Output:
	// test.Equal(t, "Hello World\nG'day World\n", "Hello World\nG'day Mate"):
	// --- got
	// +++ want
	// @@ -1,3 +1,2 @@
	//  Hello World
	// -G'day World
	// -
	// +G'day Mate
}

func ExampleNotEqual_pass() {
	test.NotEqual(t, Abs(-1), -1)
	// Output:
	// test.NotEqual(t, Abs(-1), -1): got 1, not -1
}

func ExampleNotEqual_fail() {
	test.NotEqual(t, Abs(-1), 1)
	// Output:
	// test.NotEqual(t, Abs(-1), 1): got 1, want not 1
}

func ExampleEqualJSON_pass() {
	test.EqualJSON(
		t,
		[]byte(`{"key":"value1","key":"value2","key3":3}`),
		[]byte(` {
				"key":"value2",
				"key":"value1",
				"key3": 3
			}`),
	)
	// Output:
	// test.EqualJSON(: got {
	//   "key": "value1",
	//   "key": "value2",
	//   "key3": 3
	// }
}

func ExampleEqualJSON_fail() {
	test.EqualJSON(
		t,
		[]byte(`{"key":"value1","key3":3}`),
		[]byte(` {
				"key3": 3
				"key":"value2",
			}`),
	)
	// Output:
	// test.EqualJSON(:
	// --- got
	// +++ want
	// @@ -1,5 +1,5 @@
	//  {
	// -  "key": "value1",
	// +  "key": "value2",
	//    "key3": 3
	//  }
}
