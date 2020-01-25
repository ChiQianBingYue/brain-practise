package linkedlist

import (
	"strconv"
	"strings"
)

func (l *ListNode) print() string {
	var sb strings.Builder
	for l != nil {
		s := strconv.Itoa(l.Val)
		if s == "" {
			s = "nil"
		}
		sb.WriteString(s)
		sb.WriteString(",")
		l = l.Next
	}
	return sb.String()
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
