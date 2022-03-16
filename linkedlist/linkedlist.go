package linkedlist

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (l *LinkedList[T]) add(data T) {
	node := &Node[T]{data: data}
	if l.head == nil {
		l.head = node
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (l *LinkedList[T]) print() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// func main() {
// 	ll := new(LinkedList[float32])
// 	ll.add(1.01)
// 	ll.add(2.02)
// 	ll.add(3.03)
// 	ll.print()
// }
