package linkedList

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	sizeA := getSize(headA)
	sizeB := getSize(headB)
	var more, less *ListNode
	var n int
	if sizeA < sizeB {
		less = &ListNode{Next: headA}
		more = &ListNode{Next: headB}
		n = sizeB - sizeA
	} else {
		more = &ListNode{Next: headA}
		less = &ListNode{Next: headB}
		n = sizeA - sizeB
	}
	nodeMore := more
	nodeLess := less
	for i := 0; i < n; i++ {
		nodeMore = nodeMore.Next
	}
	for nodeMore != nodeLess {
		nodeMore = nodeMore.Next
		nodeLess = nodeLess.Next
	}
	return nodeMore
}

func getSize(head *ListNode) int {
	size := 0
	node := head
	for node != nil {
		size++
		node = node.Next
	}
	return size
}
