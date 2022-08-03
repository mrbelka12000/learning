package linkedlist

//Find node from linked list by his value
func (l *LinkedList) Find(value interface{}) (*Node, int) {

	head := l.Head
	pos := 0

	//iterates over linked list
	for head != nil {

		//if values are equal we find it
		if head.Value == value {
			return head, pos
		}

		//switch head and to next
		head = head.Next
		pos++
	}

	return nil, 0
}
