package linkedlist

func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{}
	node := head
	for node != nil {
		target := node
		node = node.Next
		target.Next = newHead.Next
		newHead.Next = target
	}
	return newHead.Next
}
