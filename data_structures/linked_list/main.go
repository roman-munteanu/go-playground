package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func (n Node) String() string {
	return fmt.Sprintf("{%d}", n.data)
}

type LinkedList struct {
	head *Node
}

func (ls *LinkedList) insertAfterLastNode(data int) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if ls.head == nil {
		ls.head = newNode
		return
	}

	current := ls.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (ls *LinkedList) print() {
	current := ls.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (ls *LinkedList) insertBeforeHead(data int) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if ls.head == nil {
		ls.head = newNode
		return
	}

	newNode.next = ls.head
	ls.head = newNode
}

func (ls *LinkedList) insertAfterValue(val int, data int) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if ls.head == nil {
		ls.head = newNode
		return
	}

	current := ls.head
	for current != nil {
		if current.data == val {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}
}

func (ls *LinkedList) insertBeforeValue(val int, data int) {
	if ls.head == nil {
		return
	}

	newNode := &Node{
		data: data,
		next: nil,
	}

	if ls.head.data == val {
		newNode.next = ls.head
		ls.head = newNode
	}

	current := ls.head
	for current.next != nil {
		if current.next.data == val {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}
}

func (ls *LinkedList) deleteHead() {
	if ls.head == nil {
		return
	}

	ls.head = ls.head.next
}

func (ls *LinkedList) deleteLastNode() {
	if ls.head == nil {
		return
	}

	if ls.head.next == nil {
		ls.head = nil
		return
	}

	current := ls.head
	for current.next != nil {
		if current.next.next == nil {
			current.next = nil
			return
		}
		current = current.next
	}
}

func (ls *LinkedList) deleteAfterValue(val int) {
	current := ls.head
	for current != nil {
		if current.data == val && current.next != nil {
			current.next = current.next.next
			return
		}

		current = current.next
	}
}

func (ls *LinkedList) size() int {
	count := 0
	current := ls.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (ls *LinkedList) findNodeAt(idx int) *Node {
	current := ls.head
	for c := 0; c < idx; c++ {
		current = current.next
	}

	return current
}

func main() {
	ls := LinkedList{}
	ls.insertAfterLastNode(10)
	ls.insertAfterLastNode(20)
	ls.insertAfterLastNode(30)
	ls.insertBeforeHead(40)
	ls.insertAfterLastNode(50)

	ls.insertAfterValue(20, 2)
	ls.insertAfterValue(30, 3)

	ls.insertBeforeValue(20, 22)
	ls.insertBeforeValue(50, 55)

	// ls.deleteHead()
	// ls.deleteLastNode()
	// ls.deleteLastNode()

	ls.deleteAfterValue(30)

	ls.print()
	// fmt.Println(ls.head.data)
	fmt.Printf("size: %d\n", ls.size())

	fmt.Printf("node at index: %s\n", ls.findNodeAt(3))
}
