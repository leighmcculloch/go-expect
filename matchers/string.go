package matchers

func StringEq(expected string) func(actual string) bool {
	return func(actual string) bool {
		return actual == expected
	}
}
