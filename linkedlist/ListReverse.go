package linkedlist

func (l *LinkedList) Reverse() {

	node := l.Head

	var prev *Node
	var next *Node

	for node != nil {

		//saves next node for iterating
		next = node.Next

		//replaces next node to nil at first iteration,
		//will subsequently be the previous value from the current position
		node.Next = prev

		//sets prev to current node
		prev = node

		//move to the next node
		node = next

		//simple solution but hard for understanding
		//node, prev, node.Next = node.Next, node, prev
	}

	//because head is nil at this moment should set last value from linked list
	l.Head = prev
}
