package linkedlist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resDummy := &ListNode{}
	cur := resDummy
	leading := 0
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}
		sum := v1 + v2 + leading
		cur.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		leading = sum / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if leading == 1 {
		cur.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return resDummy.Next
}
