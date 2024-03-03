package stackqueue

func isValid(s string) bool {
	// 栈
	stack := make(Stack, 0, len(s))
	for i := 0; i < len(s); i++ {
		t := s[i]
		if t == '(' || t == '{' || t == '[' {
			push(&stack, t)
		} else if t == ')' {
			if b := pop(&stack); b != '(' {
				return false
			}
		} else if t == '}' {
			if b := pop(&stack); b != '{' {
				return false
			}
		} else if t == ']' {
			if b := pop(&stack); b != '[' {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

type Stack []byte

func push(s *Stack, b byte) {
	*s = append(*s, b)
}
func pop(s *Stack) byte {
	if len(*s) == 0 {
		return '0'
	}
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}
