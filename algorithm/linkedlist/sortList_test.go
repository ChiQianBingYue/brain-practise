package linkedlist

import (
	"fmt"
	"testing"
)

func Test_sortList(t *testing.T) {
	head := sliceToListNode([]int{4, 2, 1, 3})
	want := sliceToListNode([]int{1, 2, 3, 4})
	res := sortList(head)
	fmt.Println(res.print())
	if !want.equal(res) {
		t.Errorf("get %s, want %s", res.print(), want.print())
	}
}
