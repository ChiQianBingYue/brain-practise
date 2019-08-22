package linkedlist

// ListNode 是链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) equal(b *ListNode) bool {
	li, bi := l, b

	for li != nil && bi != nil {
		if li.Val != bi.Val {
			return false
		}
		li, bi = li.Next, bi.Next
	}
	if li != nil || bi != nil {
		return false
	}
	return true
}

func (l *ListNode) getTail() *ListNode {
	cur := l
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	return cur
}

func sliceToListNode(slice []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, n := range slice {
		cur.Next = &ListNode{}
		cur = cur.Next
		cur.Val = n
	}
	return head.Next
}
