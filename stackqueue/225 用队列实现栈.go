package stackqueue

type node struct {
	Val  int
	Next *node
}

type queue struct {
	head *node // 真实头节点
	tail *node // 尾节点
	size int
}

func (q *queue) enqueue(x int) {
	if q.size == 0 {
		q.head = &node{Val: x}
		q.tail = q.head
		q.size++
		return
	}
	q.tail.Next = &node{Val: x}
	q.tail = q.tail.Next
	q.size++
	return
}
func (q *queue) dequeue() int {
	if q.size == 0 {
		return -1
	}
	n := q.head
	q.head = q.head.Next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return n.Val
}
func (q *queue) peek() int {
	if q.size == 0 {
		return -1
	}
	return q.head.Val
}

type MyStack struct {
	q queue
}

func Constructor2() MyStack {
	return MyStack{q: queue{}}
}

func (this *MyStack) Push(x int) {
	this.q.enqueue(x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < this.q.size-1; i++ {
		this.q.enqueue(this.q.dequeue())
	}
	return this.q.dequeue()
}

func (this *MyStack) Top() int {
	for i := 0; i < this.q.size-1; i++ {
		this.q.enqueue(this.q.dequeue())
	}
	v := this.q.peek()
	this.q.enqueue(this.q.dequeue())
	return v
}

func (this *MyStack) Empty() bool {
	if this.q.size == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor2();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
