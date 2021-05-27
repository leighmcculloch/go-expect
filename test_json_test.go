package test_test

import (
	"testing"

	"4d63.com/test"
)

func TestEqJSON(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		ft := &fakeT{}
		b := test.EqJSON(
			ft,
			[]byte(`{"key":"valee","key":"value","key3":3}`),
			[]byte(` {
				"key":"value",
				"key":"valee",
				"key3": 3
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
		b := test.EqJSON(
			ft,
			[]byte(`{"key":"v alue","key3":3}`),
			[]byte(` {
				"key":"value",
				"key3": 3
			}`),
		)
		wantErr := `b := test.EqJSON(:
--- Want
+++ Got
@@ -1,5 +1,5 @@
 {
-  "key": "value",
+  "key": "v alue",
   "key3": 3
 }
 
`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Fatalf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
}
