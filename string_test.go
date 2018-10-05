package expect_test

import (
	"strings"
	"testing"

	"4d63.com/expect"
)

func TestExpectStringToEq(t *testing.T) {
	ft := &fakeT{}

	expect := expect.In(ft)

	expect.Str("hello").ToEq("world")
	expect.Str("hello").ToEq("hello")

	if g, w := len(ft.LogCalls), 1; g != w {
		t.Errorf("Log called %d times, want %d", g, w)
	} else {
		if g, w := ft.LogCalls[0], `got "hello"`; g != w {
			t.Errorf("Log call with %q, want %q", g, w)
		}
	}
	if g, w := len(ft.ErrorCalls), 1; g != w {
		t.Errorf("Error called %d times, want %d", g, w)
	} else {
		if g, w := ft.ErrorCalls[0], `got "hello", want "world"`; g != w {
			t.Errorf("Error call with %q, want %q", g, w)
		}
	}
	if g, w := len(ft.FatalCalls), 0; g != w {
		t.Errorf("Fatal called %d times, want %d", g, w)
	}
}

func TestExpectStringTo(t *testing.T) {
	ft := &fakeT{}

	expect := expect.In(ft)

	ieq := func(expected string) func(s string) bool {
		return func(s string) bool {
			return strings.ToLower(s) == strings.ToLower(expected)
		}
	}

	expect.Str("hello").To(ieq("hELLo"))
	expect.Str("hello").To(ieq("hELio"))
	expect.Str("hello").NotTo(ieq("hELLo"))
	expect.Str("hello").NotTo(ieq("hELio"))

	if g, w := len(ft.LogCalls), 2; g != w {
		t.Errorf("Log called %d times, want %d", g, w)
	}
	if g, w := len(ft.ErrorCalls), 2; g != w {
		t.Errorf("Error called %d times, want %d", g, w)
	}
	if g, w := len(ft.FatalCalls), 0; g != w {
		t.Errorf("Fatal called %d times, want %d", g, w)
	}
}
