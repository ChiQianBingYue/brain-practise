package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// use merge sort
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left := head
	right := findMidAndCut(head)
	return merge(sortList(left), sortList(right))
}

func findMidAndCut(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return mid
}

func merge(left, right *ListNode) *ListNode {
	res := &ListNode{}
	cur := res

	for left != nil && right != nil {
		if left.Val > right.Val {
			cur.Next = right
			right = right.Next
		} else {
			cur.Next = left
			left = left.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	} else if right != nil {
		cur.Next = right
	}
	return res.Next
}
