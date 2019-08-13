package linkedlist

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	l1 := sliceToListNode([]int{1, 2, 4})
	l2 := sliceToListNode([]int{1, 3, 4})
	res := sliceToListNode([]int{1, 1, 2, 3, 4, 4})
	if got := mergeTwoLists(l1, l2); !got.equal(res) {
		t.Errorf("reverseList() = %s, want %s", got.print(), res.print())
	}
}
