package test

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

// displayDumpDiff returns if a diff should be displayed as a spew dump when
// comparing the values.
func displayDumpDiff(v1, v2 interface{}) bool {
	// If the types are different, display a dump diff.
	t1 := reflect.TypeOf(v1)
	t2 := reflect.TypeOf(v2)
	if t1 != t2 {
		return true
	}

	// If the type of the values is a simple primitive type, don't display a
	// dump diff.
	switch t1.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		return false
	}
	return true
}

// Equal compares got to want and reports an error to tb if they are not equal.
// Returns true if equal.
func Equal(tb testing.TB, got, want interface{}) bool {
	tb.Helper()

	eq := cmp.Equal(got, want)
	if eq {
		tb.Logf("%s: got %+v", caller(), got)
		return eq
	}

	gotStr, gotStrOK := got.(string)
	wantStr, wantStrOK := want.(string)
	if gotStrOK && wantStrOK {
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(gotStr),
			B:        difflib.SplitLines(wantStr),
			FromFile: "got",
			ToFile:   "want",
			Context:  3,
		}
		text, _ := difflib.GetUnifiedDiffString(diff)
		tb.Errorf("%s:\n%s", caller(), text)
		return eq
	}
	if displayDumpDiff(got, want) {
		spew := spew.ConfigState{
			Indent:                  " ",
			DisableMethods:          true,
			DisablePointerAddresses: true,
			DisableCapacities:       true,
			SortKeys:                true,
			SpewKeys:                true,
		}
		gotDump := spew.Sdump(got)
		wantDump := spew.Sdump(want)
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(gotDump),
			B:        difflib.SplitLines(wantDump),
			FromFile: "got",
			ToFile:   "want",
			Context:  3,
		}
		text, _ := difflib.GetUnifiedDiffString(diff)
		tb.Errorf("%s:\n%s", caller(), text)
		return eq
	}

	tb.Errorf("%s: got %+v, want %+v", caller(), got, want)
	return eq
}

// NotEqual compares got to want and reports an error to tb if they are equal.
// Returns true if not equal.
func NotEqual(tb testing.TB, got, notWant interface{}) bool {
	tb.Helper()
	notEqual := !cmp.Equal(got, notWant)
	if notEqual {
		tb.Logf("%s: got %+v, not %+v", caller(), got, notWant)
	} else {
		tb.Errorf("%s: got %+v, want not %+v", caller(), got, notWant)
	}
	return notEqual
}

func caller() string {
	const maxCallDepth = 10
	callers := [maxCallDepth]uintptr{}
	count := runtime.Callers(0, callers[:])
	frames := runtime.CallersFrames(callers[:count])
	frame := (*runtime.Frame)(nil)
	for {
		nextFrame, more := frames.Next()
		if !more {
			return "_"
		}
		if strings.HasPrefix(nextFrame.Function, "runtime.") {
			continue
		}
		if strings.HasPrefix(nextFrame.Function, "4d63.com/test.") {
			continue
		}
		frame = &nextFrame
		break
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
