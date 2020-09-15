package main

import "fmt"

// List struct which has the information of the node
type List struct {
	head, tail *Node
}

// Node struct which has the information to the next one
type Node struct {
	value int
	next  *Node
}

// Pointing to the first node or element
func (l *List) first() *Node {
	return l.head
}
func (n *Node) second() *Node {
	return n.next
}

func (l *List) move(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func main() {
	list := &List{}
	for i := 1; i <= 5; i++ {
		list.move(i)
	}

	n := list.first()
	fmt.Println(n.value)

	n = n.second()
	fmt.Println(n.value)

	n = n.second()
	fmt.Println(n.value)

}