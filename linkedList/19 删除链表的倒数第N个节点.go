package linkedList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 获取长度
	newHead := &ListNode{Next: head}
	node := newHead.Next
	size := 0
	for node != nil {
		size++
		node = node.Next
	}

	node = newHead
	for i := 0; i < size-n; i++ {
		node = node.Next
	}
	if node.Next.Next != nil {
		node.Next = node.Next.Next
	} else {
		node.Next = nil
	}
	return newHead.Next
}
