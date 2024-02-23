package linkedList

type MyLinkedList struct {
	Head *Node
	Size int
}

type Node struct {
	Prev *Node
	Next *Node
	Val  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{Head: &Node{}}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	node := this.Head
	for i := 0; i < index+1; i++ {
		node = node.Next
	}
	return node.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{Val: val}
	headNext := this.Head.Next
	head := this.Head
	head.Next = newNode
	newNode.Next = headNext
	newNode.Prev = head
	if headNext != nil {
		headNext.Prev = newNode
	}
	this.Size++
}

// 1
// 1 3

func (this *MyLinkedList) AddAtTail(val int) {
	node := this.Head
	for i := 0; i < this.Size; i++ {
		node = node.Next
	}
	newNode := &Node{Val: val}
	node.Next = newNode
	newNode.Prev = node
	this.Size++
}

// 1 3

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}
	node := this.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	newNode := &Node{Val: val}
	nodeNext := node.Next
	node.Next = newNode
	newNode.Prev = node
	newNode.Next = nodeNext
	nodeNext.Prev = newNode
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	node := this.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	if node.Next != nil {
		node.Next.Prev = node
	}
	this.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
