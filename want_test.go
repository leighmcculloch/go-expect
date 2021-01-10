package want_test

import (
	"io"
	"testing"

	"4d63.com/want"
)

func TestEq(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		ft := &fakeT{}
		b := want.Eq(ft, "a", "a")
		if len(ft.ErrorCalls) != 0 {
			t.Errorf("got %+q, want 0 errors", ft.ErrorCalls)
		}
		if !b {
			t.Errorf("got %v, want true", b)
		}
	})
	t.Run("fail without diff", func(t *testing.T) {
		ft := &fakeT{}
		b := want.Eq(ft, 0, 1)
		wantErr := `b := want.Eq(ft, 0, 1): got 0, want 1`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Fatalf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
	t.Run("fail with string diff when comparing strings", func(t *testing.T) {
		ft := &fakeT{}
		b := want.Eq(ft, "a\nb\nc\nd\ne\nf\ng", "a\nz\nc\nd\ne\nf\ng")
		wantErr := `b := want.Eq(ft, "a\nb\nc\nd\ne\nf\ng", "a\nz\nc\nd\ne\nf\ng"):
--- Want
+++ Got
@@ -1,5 +1,5 @@
 a
-z
+b
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
		b := want.Eq(ft, value{"A"}, value{"B"})
		wantErr := `b := want.Eq(ft, value{"A"}, value{"B"}):
--- Want
+++ Got
@@ -1,4 +1,4 @@
 (want_test.value) {
- Name: (string) (len=1) "B"
+ Name: (string) (len=1) "A"
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
		b := want.NotEq(ft, "a", "b")
		if len(ft.ErrorCalls) != 0 {
			t.Fatalf("got %+q, want 0 errors", ft.ErrorCalls)
		}
		if !b {
			t.Errorf("got %v, want true", b)
		}
	})
	t.Run("fail", func(t *testing.T) {
		ft := &fakeT{}
		b := want.NotEq(ft, "a", "a")
		wantErr := `b := want.NotEq(ft, "a", "a"): got a, want not a`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Fatalf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
}

func TestNil(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		ft := &fakeT{}
		var err error
		b := want.Nil(ft, err)
		if len(ft.ErrorCalls) != 0 {
			t.Errorf("got %+q, want 0 errors", ft.ErrorCalls)
		}
		if !b {
			t.Errorf("got %v, want true", b)
		}
	})
	t.Run("fail", func(t *testing.T) {
		ft := &fakeT{}
		var err = error(io.EOF)
		b := want.Nil(ft, err)
		wantErr := `b := want.Nil(ft, err): got EOF, want <nil>`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Errorf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
}

func TestNotNil(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		ft := &fakeT{}
		var err = error(io.EOF)
		b := want.NotNil(ft, err)
		if len(ft.ErrorCalls) != 0 {
			t.Errorf("got %+q, want 0 errors", ft.ErrorCalls)
		}
		if !b {
			t.Errorf("got %v, want true", b)
		}
	})
	t.Run("fail", func(t *testing.T) {
		ft := &fakeT{}
		var err error
		b := want.NotNil(ft, err)
		wantErr := `b := want.NotNil(ft, err): got <nil>, want not <nil>`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Errorf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
}

func TestTrue(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		ft := &fakeT{}
		b := want.True(ft, true)
		if len(ft.ErrorCalls) != 0 {
			t.Errorf("got %+q, want 0 errors", ft.ErrorCalls)
		}
		if !b {
			t.Errorf("got %v, want true", b)
		}
	})
	t.Run("fail", func(t *testing.T) {
		ft := &fakeT{}
		b := want.True(ft, false)
		wantErr := `b := want.True(ft, false): got false, want true`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Errorf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
}

func TestFalse(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		ft := &fakeT{}
		b := want.False(ft, false)
		if len(ft.ErrorCalls) != 0 {
			t.Errorf("got %+q, want 0 errors", ft.ErrorCalls)
		}
		if !b {
			t.Errorf("got %v, want true", b)
		}
	})
	t.Run("fail", func(t *testing.T) {
		ft := &fakeT{}
		b := want.False(ft, true)
		wantErr := `b := want.False(ft, true): got true, want false`
		if len(ft.ErrorCalls) != 1 || ft.ErrorCalls[0] != wantErr {
			t.Errorf("got %+q, want 1 error %q", ft.ErrorCalls, wantErr)
		}
		if b {
			t.Errorf("got %v, want false", b)
		}
	})
}
