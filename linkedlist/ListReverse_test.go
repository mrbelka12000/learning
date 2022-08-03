package linkedlist

import "testing"

func TestListReverse(t *testing.T) {
	l := &LinkedList{}

	l.ListPushBack(1)
	l.ListPushBack(2)
	l.ListPushBack(5)
	l.ListPushBack(3)

	l.Reverse()

	l.Display()

}
