package stringproblem

func isValid(s string) bool {
	stack := []rune{}
	for _, r := range s {
		var l rune
		switch r {
		case rune(']'):
			if l, stack = popLast(stack); l != rune('[') {
				return false
			}
		case rune('}'):
			if l, stack = popLast(stack); l != rune('{') {
				return false
			}
		case rune(')'):
			if l, stack = popLast(stack); l != rune('(') {
				return false
			}
		default:
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}

func popLast(s []rune) (rune, []rune) {
	var res rune
	if len(s) > 0 {
		res = s[len(s)-1]
		s = s[0 : len(s)-1]
	}
	return res, s
}
