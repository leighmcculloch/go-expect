package want

import "testing"

var def = Want{}

// Eq compares got to want and reports an error to tb if they are not equal.
// Returns true if equal.
func Eq(tb testing.TB, got, want interface{}) bool {
	tb.Helper()
	return def.Eq(tb, got, want)
}

// NotEq compares got to want and reports an error to tb if they are equal.
// Returns true if not equal.
func NotEq(tb testing.TB, got, notWant interface{}) bool {
	tb.Helper()
	return def.NotEq(tb, got, notWant)
}

// Nil checks if got is nil and reports an error to tb if it is not nil.
// Returns true if nil.
func Nil(tb testing.TB, got interface{}) bool {
	tb.Helper()
	return def.Nil(tb, got)
}

// NotNil checks if got is not nil and reports an error to tb if it is nil.
// Returns true if not nil.
func NotNil(tb testing.TB, got interface{}) bool {
	tb.Helper()
	return def.NotNil(tb, got)
}

// True checks if got is true and reports an error to tb if it is not true.
// Returns true if true.
func True(tb testing.TB, got bool) bool {
	tb.Helper()
	return def.True(tb, got)
}

// False checks if got is false and reports an error to tb if it is not false.
// Returns true if false.
func False(tb testing.TB, got bool) bool {
	tb.Helper()
	return def.False(tb, got)
}

// EqJSON compares got to want and reports an error to tb if they are not equal.
// JSON is first formatted consistently and keys are sorted before comparing.
// Returns true if logically equal.
func EqJSON(tb testing.TB, got, want []byte) bool {
	tb.Helper()
	return def.EqJSON(tb, got, want)
}
