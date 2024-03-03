package stackqueue

import "strconv"

func evalRPN(tokens []string) int {
	stack := make(Stack2, 0, len(tokens))
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if t == "+" {
			e1, _ := pop2(&stack)
			e2, _ := pop2(&stack)
			push2(&stack, e1+e2)
		} else if t == "-" {
			e1, _ := pop2(&stack)
			e2, _ := pop2(&stack)
			push2(&stack, e2-e1)
		} else if t == "*" {
			e1, _ := pop2(&stack)
			e2, _ := pop2(&stack)
			push2(&stack, e2*e1)
		} else if t == "/" {
			e1, _ := pop2(&stack)
			e2, _ := pop2(&stack)
			push2(&stack, e2/e1)
		} else {
			i, _ := strconv.Atoi(t)
			push2(&stack, i)
		}
	}
	return stack[0]
}

type Stack2 []int

func push2(s *Stack2, str int) {
	*s = append(*s, str)
}
func pop2(s *Stack2) (int, bool) {
	size := len(*s)
	if size == 0 {
		return 0, false
	}
	res := (*s)[size-1]
	*s = (*s)[:size-1]
	return res, true
}
