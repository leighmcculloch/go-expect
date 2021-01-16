package want_test

import (
	"testing"

	"4d63.com/want"
	"github.com/tidwall/pretty"
)

func TestEqJSON_PRETTY(t *testing.T) {
	opt := pretty.Options{
		Indent:   "  ",
		SortKeys: true,
	}

	{
		in := `{"key":"value","key":"value2","key3":3}`
		out := pretty.PrettyOptions([]byte(in), &opt)
		t.Log(string(out))
	}

	{
		in := ` {
				"key":"value2",
				"key":"value",
				"key3": 3
			}`
		out := pretty.PrettyOptions([]byte(in), &opt)
		t.Log(string(out))
	}
}

func TestEqJSON(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		ft := &fakeT{}
		// TODO: pretty seems to not be consistently sorting keys, prove that and open an issue
		b := want.EqJSON(
			ft,
			[]byte(`{"key":"value","key":"value2","key3":3}`),
			[]byte(` {
				"key":"value2",
				"key":"value"
				"key3": 3,
			}`),
		)
		if len(ft.ErrorCalls) != 0 {
			t.Errorf("got %+q, want 0 errors", ft.ErrorCalls)
		}
		if !b {
			t.Errorf("got %v, want true", b)
		}
	})
	t.Run("fail with string diff", func(t *testing.T) {
		ft := &fakeT{}
		b := want.EqJSON(
			ft,
			[]byte(`{"key":"v alue","key":"value2","key3":3}`),
			[]byte(` {
				"key":"value",
				"key":"value2",
				"key3": 3
			}`),
		)
		wantErr := `b := want.EqJSON(:
--- Want
+++ Got
@@ -1,5 +1,5 @@
 {
-  "key": "value",
+  "key": "v alue",
   "key": "value2",
   "key3": 3
 }
`
		want.Eq(t, ft.ErrorCalls[0], wantErr)
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Fatalf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
}
