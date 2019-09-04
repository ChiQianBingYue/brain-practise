package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lA := getNodeLen(headA)
	lB := getNodeLen(headB)
	if lA > lB {
		headA = getNthNode(headA, lA-lB)
	} else {
		headB = getNthNode(headB, lB-lA)
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func getNthNode(head *ListNode, nth int) *ListNode {
	count := 0
	for count < nth && head != nil {
		head = head.Next
		count++
	}
	return head
}

func getNodeLen(head *ListNode) int {
	res := 0
	for head != nil {
		head = head.Next
		res++
	}
	return res
}
