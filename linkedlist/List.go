package linkedlist

type LinkedList struct {
	Len  int
	Head *Node
	Tail *Node
}

type Node struct {
	Next  *Node
	Value interface{}
}
