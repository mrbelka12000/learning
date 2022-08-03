package linkedlist

import "testing"

func TestListPushBack(t *testing.T) {
	l := &LinkedList{}

	l.ListPushBack("beka")
	l.ListPushBack("teka")
	l.ListPushBack(11)

	l.Display()
}
