package linkedlist

import "testing"

func Test_getIntersectionNode(t *testing.T) {
	l1 := sliceToListNode([]int{1, 4, 5})
	l2 := sliceToListNode([]int{1, 3, 4})
	l3 := sliceToListNode([]int{2, 6})
	l1.getTail().Next = l3
	l2.getTail().Next = l3

	res := getIntersectionNode(l1, l2)
	if res != l3 {
		t.Errorf("getIntersectionNode() = %s, want %s", res.print(), l3.print())
	}
}
