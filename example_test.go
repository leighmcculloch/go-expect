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

func ExampleT_Equal_pass() {
	t := test.T{t}
	t.Equal(Abs(-1), 1)
	// Output:
	// t.Equal(Abs(-1), 1): got 1
}

func ExampleT_Equal_fail() {
	tt := test.New(t)
	tt.Equal(Abs(-1), 0)
	// Output:
	// tt.Equal(Abs(-1), 0): got 1, want 0
}

func ExampleT_Equal_failComparingDifferentTypes() {
	tt := test.New(t)
	tt.Equal(1, 1.0)
	// Output:
	// tt.Equal(1, 1.0):
	// --- got
	// +++ want
	// @@ -1,2 +1,2 @@
	// -(int) 1
	// +(float64) 1
}

func ExampleT_Equal_failDiff() {
	tt := test.New(t)
	tt.Equal("Hello World\nG'day World\n", "Hello World\nG'day Mate")
	// Output:
	// tt.Equal("Hello World\nG'day World\n", "Hello World\nG'day Mate"):
	// --- got
	// +++ want
	// @@ -1,3 +1,2 @@
	//  Hello World
	// -G'day World
	// -
	// +G'day Mate
}

func ExampleT_NotEqual_pass() {
	tt := test.New(t)
	tt.NotEqual(Abs(-1), -1)
	// Output:
	// tt.NotEqual(Abs(-1), -1): got 1, not -1
}

func ExampleT_NotEqual_fail() {
	tt := test.New(t)
	tt.NotEqual(Abs(-1), 1)
	// Output:
	// tt.NotEqual(Abs(-1), 1): got 1, want not 1
}

func ExampleT_EqualJSON_pass() {
	tt := test.New(t)
	tt.EqualJSON(
		[]byte(`{"key":"value1","key":"value2","key3":3}`),
		[]byte(` {
				"key":"value2",
				"key":"value1",
				"key3": 3
			}`),
	)
	// Output:
	// tt.EqualJSON(: got {
	//   "key": "value1",
	//   "key": "value2",
	//   "key3": 3
	// }
}

func ExampleT_EqualJSON_fail() {
	tt := test.New(t)
	tt.EqualJSON(
		[]byte(`{"key":"value1","key3":3}`),
		[]byte(` {
				"key3": 3
				"key":"value2",
			}`),
	)
	// Output:
	// tt.EqualJSON(:
	// --- got
	// +++ want
	// @@ -1,5 +1,5 @@
	//  {
	// -  "key": "value1",
	// +  "key": "value2",
	//    "key3": 3
	//  }
}
