package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

func reverseList(head *ListNode) *ListNode {
	// _, res := reverseListRecursive(head)
	return reverseListIterate(head)
}

func reverseListRecursive(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	if head.Next == nil {
		return head, head
	}
	next, tail := reverseListRecursive(head.Next)
	next.Next = head
	head.Next = nil
	return head, tail
}

func reverseListIterate(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, cur, next := head, head.Next, head.Next.Next
	pre.Next = nil
	cur.Next = pre
	for next != nil {
		pre, cur, next = cur, next, next.Next
		cur.Next = pre
	}
	return cur
}
