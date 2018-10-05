package expect

import "testing"

func In(tb testing.TB) Expect {
	return Expect{tb: tb}
}

type Expect struct {
	tb  testing.TB
	not bool
}

func (x Expect) eval(b bool) bool {
	return (!x.not && b) || (x.not && !b)
}

func (x Expect) In(tb testing.TB) Expect {
	x.tb = tb
	return x
}

func (x Expect) Str(s string) String {
	return String{x: x, s: s}
}
