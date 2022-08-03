package linkedlist

//ListPushBack adds element to the end of list
func (l *LinkedList) ListPushBack(value interface{}) {
	node := &Node{Value: value}

	//if linked list empty
	if l.Head == nil {
		l.Tail = node
		l.Head = node
	} else {
		// adds next to tail in list and sets tail
		l.Tail.Next = node
		l.Tail = node
	}
	l.Len++
}
