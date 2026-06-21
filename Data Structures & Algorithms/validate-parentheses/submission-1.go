func isValid(s string) bool {
    if len(s)%2 != 0 {
        return false
    }
    stack := []rune{}
    for _, c := range s {
        if c == '(' || c == '[' || c == '{' {
			push(&stack, c)
		} else if c == ')' || c == ']' || c == '}' {
            if isEmpty(&stack) {
				return false
			}
			prev := peek(&stack)
			if prev == '(' && c == ')' || prev == '[' && c == ']' || prev == '{' && c == '}' {
				pop(&stack)
			} else {
				return false
			}
		}
    }
    return isEmpty(&stack)
}

func push(stack *[]rune, val rune) {
	*stack = append(*stack, val)
}

func pop(stack *[]rune) rune {
	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return top
}

func peek(stack *[]rune) rune {
	return (*stack)[len(*stack)-1]
}

func isEmpty(stack *[]rune) bool {
	return len(*stack) == 0
}