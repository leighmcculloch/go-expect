package want

import (
	"io/ioutil"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/pmezard/go-difflib/difflib"
)

// A Want is a set of options for configuring the behavior of the library. Its
// zero value (Want{}) is usable and is equivalent to invoking the package
// functions.
type Want struct{}

// displayStringDiff returns if a diff should be displayed as a simple comparison
// of two strings when comparing the value.
func displayStringDiff(v interface{}) bool {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.String:
		return true
	}
	return false
}

// displayDumpDiff returns if a diff should be displayed as a spew dump when
// comparing the value.
func displayDumpDiff(v interface{}) bool {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		return false
	}
	return true
}

// Eq compares got to want and reports an error to tb if they are not equal.
// Returns true if equal.
func (w *Want) Eq(tb testing.TB, got, want interface{}) bool {
	tb.Helper()

	eq := cmp.Equal(got, want)
	if eq {
		tb.Logf("%s: got %+v", w.caller(), got)
		return eq
	}

	if displayStringDiff(got) || displayStringDiff(want) {
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(want.(string)),
			B:        difflib.SplitLines(got.(string)),
			FromFile: "Want",
			ToFile:   "Got",
			Context:  3,
		}
		text, _ := difflib.GetUnifiedDiffString(diff)
		tb.Errorf("%s:\n%s", w.caller(), text)
		return eq
	}

	if displayDumpDiff(got) || displayDumpDiff(want) {
		spew := spew.ConfigState{
			Indent:                  " ",
			DisableMethods:          true,
			DisablePointerAddresses: true,
			DisableCapacities:       true,
			SortKeys:                true,
			SpewKeys:                true,
		}
		gotS := spew.Sdump(got)
		wantS := spew.Sdump(want)
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(wantS),
			B:        difflib.SplitLines(gotS),
			FromFile: "Want",
			ToFile:   "Got",
			Context:  3,
		}
		text, _ := difflib.GetUnifiedDiffString(diff)
		tb.Errorf("%s:\n%s", w.caller(), text)
		return eq
	}

	tb.Errorf("%s: got %+v, want %+v", w.caller(), got, want)
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
		tb.Logf("%s: got %+v, not <nil>", w.caller(), got)
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
