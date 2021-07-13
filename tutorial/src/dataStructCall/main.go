package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	//list := &dataStruct.LinkedList{}
	//list.AddNode(0)
	//for i := 1; i < 10; i++ {
	//	list.AddNode(i)
	//}
	//list.PrintNodes()
	//list.RemoveNode(list.Root.Next)
	//list.PrintNodes()

	//list.RemoveNode(list.Root)
	//list.PrintNodes()

	//list.RemoveNode(list.Tail)
	//list.PrintNodes()

	//fmt.Printf("%d\n", list.Tail.Val)

	//list.PrintReverse()

	///************************************/

	//stack2 := dataStruct.NewStack()

	//for i := 1; i <= 5; i++ {
	//	stack2.Push(i)
	//}

	//fmt.Println("NewStack")
	//for !stack2.Empty() {
	//	val := stack2.Pop()
	//	fmt.Printf("%d -> ", val)
	//}
	//fmt.Println()

	//queue2 := dataStruct.NewQueue()
	//for i := 1; i <= 5; i++ {
	//	queue2.Push(i)
	//}

	//fmt.Println("NewQueue")
	//for !queue2.Empty() {
	//	val := queue2.Pop()
	//	fmt.Printf("%d -> ", val)
	//}

	///************************************/

	tree := dataStruct.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}

	tree.DFS1()
	fmt.Println()

	tree.DFS2()
	fmt.Println()

	tree.BFS()
}