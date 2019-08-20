package linkedlist

import (
	"testing"
)

func Test_mergeKListsBasic(t *testing.T) {
	l1 := sliceToListNode([]int{1, 4, 5})
	l2 := sliceToListNode([]int{1, 3, 4})
	l3 := sliceToListNode([]int{2, 6})
	list := []*ListNode{l1, l2, l3}
	res := sliceToListNode([]int{1, 1, 2, 3, 4, 4, 5, 6})
	if got := mergeKLists(list); !got.equal(res) {
		t.Errorf("reverseList() = %s, want %s", got.print(), res.print())
	}
}

func Test_mergeKListsNil(t *testing.T) {

	list := []*ListNode{}
	res := sliceToListNode([]int{})
	if got := mergeKLists(list); !got.equal(res) {
		t.Errorf("reverseList() = %s, want %s", got.print(), res.print())
	}
}
