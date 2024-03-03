package stackqueue

func removeDuplicates(s string) string {
	stack := make(Stack, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == peek(&stack) {
			pop(&stack)
		} else {
			push(&stack, s[i])
		}
	}
	return string(stack)
}

func peek(s *Stack) byte {
	if len(*s) == 0 {
		return '0'
	}
	res := (*s)[len(*s)-1]
	return res
}
