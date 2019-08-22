package linkedlist

import "testing"

func Test_hasCycle(t *testing.T) {
	head := sliceToListNode([]int{1, 2, 4})
	head.getTail().Next = head.Next
	if has := hasCycle(head); !has {
		t.Errorf("hasCycle 基础用例错误")
	}

	head = &ListNode{Val: 1}
	head.Next = head
	if has := hasCycle(head); !has {
		t.Errorf("hasCycle 单个循环 错误")
	}
	head = &ListNode{Val: 1}
	if has := hasCycle(head); has {
		t.Errorf("hasCycle 单个不循环 错误")
	}

	head = sliceToListNode([]int{1, 2})
	head.Next = head
	if has := hasCycle(head); !has {
		t.Errorf("hasCycle 2个循环 错误")
	}
	head = sliceToListNode([]int{1, 2})
	if has := hasCycle(head); has {
		t.Errorf("hasCycle 2个不循环 错误")
	}

	head = nil
	if has := hasCycle(head); has {
		t.Errorf("hasCycle nil 错误")
	}

	head = sliceToListNode([]int{1, 2, 4})
	tail := head.getTail()
	tail.Next = tail
	if has := hasCycle(head); !has {
		t.Errorf("hasCycle 复杂尾循环")
	}
}
