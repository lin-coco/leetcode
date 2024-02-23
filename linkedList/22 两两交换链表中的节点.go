package linkedList

func swapPairs(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	node := newHead
	for node.Next != nil && node.Next.Next != nil {
		one := node.Next
		two := node.Next.Next
		prev := node
		node = two.Next
		prev.Next = two
		two.Next = one
		one.Next = node
		node = one
	}
	return newHead.Next
}
