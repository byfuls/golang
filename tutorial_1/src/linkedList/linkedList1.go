package main

import "fmt"

type Node struct {
	next *Node
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
	l.tail = l.tail.next
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		node.next = nil
		return
	}

	prev := l.root
	for prev.next != node {
		prev = prev.next
	}

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		prev.next = prev.next.next
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
}

func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node

	return node
}

func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {

	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}

	return root, tail
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}
