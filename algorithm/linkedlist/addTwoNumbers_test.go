package linkedlist

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	res := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}
	if got := addTwoNumbers(l1, l2); !got.equal(res) {
		t.Errorf("reverseList() = %s, want %s", got.print(), res.print())
	}

}

func Test_addTwoNumbers2(t *testing.T) {
	l1 := &ListNode{
		Val: 5,
	}
	l2 := &ListNode{
		Val: 6,
	}
	res := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
		},
	}
	if got := addTwoNumbers(l1, l2); !got.equal(res) {
		t.Errorf("reverseList() = %s, want %s", got.print(), res.print())
	}

}
