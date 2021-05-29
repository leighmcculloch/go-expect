package test

import (
	"github.com/tidwall/pretty"
)

// EqualJSON compares got to want and reports an error to tb if they are not equal.
// JSON is first formatted consistently and keys are sorted before comparing.
// Returns true if logically equal.
func (t *T) EqualJSON(got, want []byte) bool {
	t.Helper()

	opt := pretty.Options{
		Indent:   "  ",
		SortKeys: true,
	}
	gotPretty := pretty.PrettyOptions(got, &opt)
	wantPretty := pretty.PrettyOptions(want, &opt)

	return t.Equal(string(gotPretty), string(wantPretty))
}
