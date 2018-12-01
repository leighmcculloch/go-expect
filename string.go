package expect

type String struct {
	x Expect
	s string
}

func (s String) To(f func(string) bool, formatAndArgs ...interface{}) bool {
	s.x.tb.Helper()
	r := s.x.eval(f(s.s))
	if r {
		s.x.tb.Logf("got %q", s.s)
	} else {
		if len(formatAndArgs) > 1 {
			s.x.tb.Errorf(formatAndArgs[0].(string), formatAndArgs[1:]...)
		} else if len(formatAndArgs) > 0 {
			s.x.tb.Errorf(formatAndArgs[0].(string))
		} else {
			s.x.tb.Errorf("got %q, not expected", s.s)
		}
	}
	return r
}

func (s String) NotTo(f func(string) bool, formatAndArgs ...interface{}) bool {
	s.x.tb.Helper()
	s.x.not = true
	return s.To(f, formatAndArgs...)
}
