// Package want is a package that provides equality functions for testing Go
// code. It contains only this one function and is intended to be succinct and
// focused on making the most common action in test code simple, testing
// equality expectations.
//
// A simple test function looks like this:
//
//     func TestAbs(t *testing.T) {
//         want.Eq(t, Abs(-1), 1)
//     }
//
// If the check passes, the verbose output looks like this:
//
//     --- PASS: TestAbs (0.00s)
//         test.go:2: want.Eq(t, Abs(-1), 1): got 1
//
// If the check fails, the output looks like this:
//
//     --- FAIL: TestAbs (0.00s)
//         test.go:2: want.Eq(t, Abs(-1), 1): got 0, want 1
//
// Got and want
//
// The terms got and want are used to describe what you got as a result of
// running the code, and what you want to have gotten. These terms are commonly
// found in the Go stdlib and it's own testing docs. In some other testing
// libraries they are sometimes referred to actual and expected.
//
// Nesting
//
// Checks can be nested using the bool return value of a prior check.
//
//     func TestAbs(t *testing.T) {
//         if want.Eq(t, Abs(-1), 1) {
//             ...
//         }
//     }
//
// Breaking early
//
// Checks can cause a test to stop at a failure using the bool return value.
//
//     func TestAbs(t *testing.T) {
//         if !want.Eq(t, Abs(-1), 1) {
//             return
//         }
//         ...
//     }
//
// Comparison
//
// Comparison of got and want is done using Google's cmp Go module:
// https://github.com/google/go-cmp/cmp
//
// Diffs
//
// Diffs can be enabled in error output of Eq comparisons. To enable diffs
// instantiate Want with DiffEnabled true and use its functions instead of the
// package functions.
//
// If the check fails, the output looks like this:
//
//     --- FAIL: TestAbs (0.00s)
//         test.go:2: want.Eq(t, Abs(-1), 1): {int}:
//                -: 0
//                +: 1
//
package want

import (
	"io/ioutil"
	"runtime"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

// A Want is a set of options for configuring the behavior of the library. Its
// zero value (Want{}) is usable and is equivalent to invoking the package
// functions Eq and NotEq.
type Want struct {
	// DiffEnabled when true enables comparison diffs in error reports.
	DiffEnabled bool
}

// Eq compares got to want and reports an error to tb if they are not equal.
// Returns true if equal.
func (w *Want) Eq(tb testing.TB, got, want interface{}) bool {
	tb.Helper()
	if w.DiffEnabled {
		d := cmp.Diff(got, want)
		eq := d == ""
		if eq {
			tb.Logf("%s: got %+v", w.caller(), got)
		} else {
			tb.Error(w.caller()+":", d)
		}
		return eq
	}
	eq := cmp.Equal(got, want)
	if eq {
		tb.Logf("%s: got %+v", w.caller(), got)
	} else {
		tb.Errorf("%s: got %+v, want %+v", w.caller(), got, want)
	}
	return eq
}

// NotEq compares got to want and reports an error to tb if they are equal.
// Returns true if not equal.
func (w *Want) NotEq(tb testing.TB, got, notWant interface{}) bool {
	tb.Helper()
	notEq := !cmp.Equal(got, notWant)
	if notEq {
		tb.Logf("%s: got %+v, not %+v", w.caller(), got, notWant)
	} else {
		tb.Errorf("%s: got %+v, want not %+v", w.caller(), got, notWant)
	}
	return notEq
}

// Nil checks if got is nil and reports an error to tb if it is not nil.
// Returns true if nil.
func (w *Want) Nil(tb testing.TB, got interface{}) bool {
	tb.Helper()
	isNil := got == nil
	if isNil {
		tb.Logf("%s: got %+v", w.caller(), got)
	} else {
		tb.Errorf("%s: got %+v, want <nil>", w.caller(), got)
	}
	return isNil
}

// NotNil checks if got is not nil and reports an error to tb if it is nil.
// Returns true if not nil.
func (w *Want) NotNil(tb testing.TB, got interface{}) bool {
	tb.Helper()
	notNil := got != nil
	if notNil {
		tb.Logf("%s: got %+v", w.caller(), got)
	} else {
		tb.Errorf("%s: got %+v, want not <nil>", w.caller(), got)
	}
	return notNil
}

// True checks if got is true and reports an error to tb if it is not true.
// Returns true if true.
func (w *Want) True(tb testing.TB, got bool) bool {
	tb.Helper()
	isTrue := got
	if isTrue {
		tb.Logf("%s: got %+v", w.caller(), got)
	} else {
		tb.Errorf("%s: got %+v, want true", w.caller(), got)
	}
	return isTrue
}

// False checks if got is false and reports an error to tb if it is not false.
// Returns true if false.
func (w *Want) False(tb testing.TB, got bool) bool {
	tb.Helper()
	isFalse := !got
	if isFalse {
		tb.Logf("%s: got %+v", w.caller(), got)
	} else {
		tb.Errorf("%s: got %+v, want false", w.caller(), got)
	}
	return isFalse
}

func (w *Want) caller() string {
	skip := 4
	if w == &def {
		skip = 5
	}
	callers := [10]uintptr{}
	count := runtime.Callers(0, callers[:])
	frames := runtime.CallersFrames(callers[:count])
	frame := (*runtime.Frame)(nil)
	for i := 0; i < skip; i++ {
		nextFrame, more := frames.Next()
		if !more {
			return "_"
		}
		frame = &nextFrame
	}
	fileBytes, err := ioutil.ReadFile(frame.File)
	if err != nil {
		return "_"
	}
	fileLines := strings.Split(string(fileBytes), "\n")
	if frame.Line >= len(fileLines) {
		return "_"
	}
	line := strings.TrimSpace(fileLines[frame.Line-1])
	return line
}
