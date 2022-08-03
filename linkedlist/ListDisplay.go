package linkedlist

import "fmt"

//Display prints all nodes of linked list
func (l *LinkedList) Display() {
	check := 0
	for l.Head != nil {
		if check == l.Len-1 {
			fmt.Printf("%v", l.Head.Value)
		} else {
			fmt.Printf("%v -> ", l.Head.Value)
		}
		l.Head = l.Head.Next
		check++
	}
	fmt.Println()
}
