package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node

	root = &Node{val: 0}

	for i := 1; i < 10; i++ {
		AddNode(root, i)
	}

	PrintNodes(root)

	root, tail = RemoveNode(root.next, root, tail)

	PrintNodes(root)
}

func AddNode(root *Node, val int) {
	var tail *Node
	tail = root
	for tail.next != nil {
		tail = tail.next
	}

	node := &Node{val: val}
	tail.next = node
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
