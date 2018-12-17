package examples

import (
	"math"
	"testing"

	"4d63.com/want"
)

func TestWantEq(t *testing.T) {
	want.Eq(t, math.Abs(-1), 1.0)
	want.Eq(t, math.Abs(-1), 1.0)
	want.Eq(t, math.Abs(-1), 2)
}

func TestWantTrue(t *testing.T) {
	want.True(t, false)
	want.True(t, true)
}

func TestWantFalse(t *testing.T) {
	want.False(t, true)
	want.False(t, false)
}
