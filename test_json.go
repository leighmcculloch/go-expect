package test

import (
	"testing"

	"github.com/tidwall/pretty"
)

// EqualJSON compares got to want and reports an error to tb if they are not equal.
// JSON is first formatted consistently and keys are sorted before comparing.
// Returns true if logically equal.
func EqualJSON(tb testing.TB, got, want []byte) bool {
	tb.Helper()

	opt := pretty.Options{
		Indent:   "  ",
		SortKeys: true,
	}
	gotPretty := pretty.PrettyOptions(got, &opt)
	wantPretty := pretty.PrettyOptions(want, &opt)

	return Equal(tb, string(gotPretty), string(wantPretty))
}
