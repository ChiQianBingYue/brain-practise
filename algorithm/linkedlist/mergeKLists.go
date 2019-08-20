package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else {
		mid := len(lists) / 2
		return mergeTwoLists(
			mergeKLists(lists[:mid]),
			mergeKLists(lists[mid:]),
		)
	}
}
