package linkedlist

func detectCycle(head *ListNode) *ListNode {
	// if head == nil || head.Next == nil {
	//     return nil
	// }
	// 快慢指针
	newHead := &ListNode{Next: head}
	left := newHead
	right := newHead
	for {
		right = right.Next
		if right == nil {
			return nil
		}
		right = right.Next
		if right == nil {
			return nil
		}
		left = left.Next
		if left == right {
			break
		}
	}
	left = newHead
	for left != right {
		left = left.Next
		right = right.Next
	}
	return left
}
