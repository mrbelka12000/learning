package linkedlist

import "testing"

func TestListPushFront(t *testing.T) {
	l := &LinkedList{}

	l.ListPushFront("beka")
	l.ListPushFront("teka")
	l.ListPushFront(11)

	l.Display()
}
