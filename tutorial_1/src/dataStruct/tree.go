package dataStruct

import "fmt"

/* search logic
- DFS
- BFS
*/

/*	DFS
				1
		2		3			4
	5	6		7	8		9 10
*/

/*	BFS
				1
		2		3			4
	5	6		7	8		9 10
*/

/* REF
- dijkstra 알고리즘
*/

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

func (t *Tree) DFS1() {
	DFS1(t.Root)
}

// DFS1 : Method 1 - recursive call
func DFS1(node *TreeNode) {
	fmt.Printf("%d -> ", node.Val)

	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}
}

// DFS2 : Method 2 - stack (+slice)
func (t *Tree) DFS2() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d -> ", last.Val)

		for i := len(last.Childs) - 1; i >= 0; i-- {
			s = append(s, last.Childs[i])
		}
	}
}

// BFS : queue
func (t *Tree) BFS() {
	queue := []*TreeNode{}
	queue = append(queue, t.Root)

	for len(queue) > 0 {
		var first *TreeNode
		first, queue = queue[0], queue[1:]

		fmt.Printf("%d -> ", first.Val)

		for i := 0; i < len(first.Childs); i++ {
			queue = append(queue, first.Childs[i])
		}
	}
}
