package linkedlist

//ListPushFront adds node to the start of list
func (l *LinkedList) ListPushFront(value interface{}) {
	node := &Node{Value: value}

	//if linked list empty
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		//sets nodes next to previous head and switch head with new node
		node.Next = l.Head
		l.Head = node
	}
	l.Len++
}
