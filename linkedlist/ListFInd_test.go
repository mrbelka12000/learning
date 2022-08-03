package linkedlist

import "testing"

func TestListFind(t *testing.T) {
	l := &LinkedList{}

	l.ListPushBack("check")
	l.ListPushFront(1234)
	l.ListPushBack("teka")

	node, pos := l.Find("check")
	if node == nil {
		t.FailNow()
	}
	t.Log(node, pos)
}
