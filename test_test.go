package test_test

import (
	"testing"

	"4d63.com/test"
)

func TestEqual(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		ft := &fakeT{}
		tt := test.New(ft)
		b := tt.Equal("a", "a")
		if len(ft.ErrorCalls) != 0 {
			t.Errorf("got %+q, want 0 errors", ft.ErrorCalls)
		}
		if !b {
			t.Errorf("got %v, want true", b)
		}
	})
	t.Run("fail without diff", func(t *testing.T) {
		ft := &fakeT{}
		tt := test.New(ft)
		b := tt.Equal(0, 1)
		wantErr := `b := tt.Equal(0, 1): got 0, want 1`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Fatalf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
	t.Run("fail with string diff when comparing strings", func(t *testing.T) {
		ft := &fakeT{}
		tt := test.New(ft)
		b := tt.Equal("a\nb\nc\nd\ne\nf\ng", "a\nz\nc\nd\ne\nf\ng")
		wantErr := `b := tt.Equal("a\nb\nc\nd\ne\nf\ng", "a\nz\nc\nd\ne\nf\ng"):
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
		tt := test.New(ft)
		type value struct {
			Name string
		}
		b := tt.Equal(value{"A"}, value{"B"})
		wantErr := `b := tt.Equal(value{"A"}, value{"B"}):
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

func TestNotEqual(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		ft := &fakeT{}
		tt := test.New(ft)
		b := tt.NotEqual("a", "b")
		if len(ft.ErrorCalls) != 0 {
			t.Fatalf("got %+q, want 0 errors", ft.ErrorCalls)
		}
		if !b {
			t.Errorf("got %v, want true", b)
		}
	})
	t.Run("fail", func(t *testing.T) {
		ft := &fakeT{}
		tt := test.New(ft)
		b := tt.NotEqual("a", "a")
		wantErr := `b := tt.NotEqual("a", "a"): got a, want not a`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Fatalf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
}
