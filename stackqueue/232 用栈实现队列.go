package stackqueue

type stack []int

func (s *stack) push(x int) {
	*s = append(*s, x)
}
func (s *stack) pop() int {
	if len(*s) == 0 {
		return -1
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
func (s *stack) peek() int {
	if len(*s) == 0 {
		return -1
	}
	return (*s)[len(*s)-1]
}

type MyQueue struct {
	S1 *stack
	S2 *stack
}

func Constructor() MyQueue {
	s1 := make(stack, 0)
	s2 := make(stack, 0)
	return MyQueue{S1: &s1, S2: &s2}
}

func (this *MyQueue) Push(x int) {
	this.S1.push(x)
}

func (this *MyQueue) Pop() int {
	if t := this.S2.pop(); t != -1 {
		return t
	}
	for {
		if t := this.S1.pop(); t != -1 {
			this.S2.push(t)
		} else {
			break
		}
	}
	return this.S2.pop()
}

func (this *MyQueue) Peek() int {
	if t := this.S2.peek(); t != -1 {
		return t
	}
	for {
		if t := this.S1.pop(); t != -1 {
			this.S2.push(t)
		} else {
			break
		}
	}
	return this.S2.peek()
}

func (this *MyQueue) Empty() bool {
	if this.S2.peek() == -1 && this.S1.peek() == -1 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor2();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
