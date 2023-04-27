package main

func solve(s string) bool {

	stack := []rune{}

	push := func(x rune) {
		stack = append(stack, x)
	}

	pop := func() rune {
		element := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return element
	}

	peek := func() rune {
		return stack[len(stack)-1]
	}

	for _, r := range s {
		switch r {
		case '(':
			push(')')
		case '[':
			push(']')
		case '{':
			push('}')
		default:
			if 0 < len(stack) && peek() == r {
				pop()
			} else {
				return false // fail fast, invalid sequence of parentheses
			}
		}
	}

	return len(stack) == 0
}
