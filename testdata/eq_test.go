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

func TestWantEqDiff(t *testing.T) {
	want := want.Want{DiffEnabled: true}
	want.Eq(t, math.Abs(-1), 2)
}
