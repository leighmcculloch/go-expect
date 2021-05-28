package test_test

import (
	"testing"

	"4d63.com/test"
)

func TestEq(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		ft := &fakeT{}
		b := test.Eq(ft, "a", "a")
		if len(ft.ErrorCalls) != 0 {
			t.Errorf("got %+q, want 0 errors", ft.ErrorCalls)
		}
		if !b {
			t.Errorf("got %v, want true", b)
		}
	})
	t.Run("fail without diff", func(t *testing.T) {
		ft := &fakeT{}
		b := test.Eq(ft, 0, 1)
		wantErr := `b := test.Eq(ft, 0, 1): got 0, want 1`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Fatalf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
	t.Run("fail with string diff when comparing strings", func(t *testing.T) {
		ft := &fakeT{}
		b := test.Eq(ft, "a\nb\nc\nd\ne\nf\ng", "a\nz\nc\nd\ne\nf\ng")
		wantErr := `b := test.Eq(ft, "a\nb\nc\nd\ne\nf\ng", "a\nz\nc\nd\ne\nf\ng"):
--- got
+++ want
@@ -1,5 +1,5 @@
 a
-b
+z
 c
 d
 e
`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Fatalf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
	t.Run("fail with dump diff when comparing structs", func(t *testing.T) {
		ft := &fakeT{}
		type value struct {
			Name string
		}
		b := test.Eq(ft, value{"A"}, value{"B"})
		wantErr := `b := test.Eq(ft, value{"A"}, value{"B"}):
--- got
+++ want
@@ -1,4 +1,4 @@
 (test_test.value) {
- Name: (string) (len=1) "A"
+ Name: (string) (len=1) "B"
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

func TestNotEq(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		ft := &fakeT{}
		b := test.NotEq(ft, "a", "b")
		if len(ft.ErrorCalls) != 0 {
			t.Fatalf("got %+q, want 0 errors", ft.ErrorCalls)
		}
		if !b {
			t.Errorf("got %v, want true", b)
		}
	})
	t.Run("fail", func(t *testing.T) {
		ft := &fakeT{}
		b := test.NotEq(ft, "a", "a")
		wantErr := `b := test.NotEq(ft, "a", "a"): got a, want not a`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Fatalf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
}
