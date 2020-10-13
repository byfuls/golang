package main

import "fmt"

/* DOUBLE LINKED LIST
put, O(1)
get, O(1)
*/

type Node struct {
	next *Node
	prev *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}

	l.tail.next = &Node{val: val}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		if l.root != nil {
			l.root.prev = nil
		}
		node.next = nil
		return
	}

	prev := node.prev

	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		node.prev = nil
		prev.next = prev.next.next
		prev.next.prev = prev
	}
	node.next = nil
}

func (l *LinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) PrintReverse() {
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}
	fmt.Printf("%d\n", node.val)
}

func main() {
	// --1
	//var root *Node
	//var tail *Node

	//root = &Node{val: 0}
	//tail = root

	//for i := 1; i < 10; i++ {
	//	tail = AddNode(tail, i)
	//}

	//PrintNodes(root)
	//root, tail = RemoveNode(root.next, root, tail)
	//PrintNodes(root)
	//root, tail = RemoveNode(root, root, tail)
	//PrintNodes(root)
	//root, tail = RemoveNode(tail, root, tail)
	//PrintNodes(root)

	// --2
	list := &LinkedList{}
	list.AddNode(0)
	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}
	list.PrintNodes()
	list.RemoveNode(list.root.next)
	list.PrintNodes()

	list.RemoveNode(list.root)
	list.PrintNodes()

	list.RemoveNode(list.tail)
	list.PrintNodes()

	fmt.Printf("%d\n", list.tail.val)

	list.PrintReverse()
}
