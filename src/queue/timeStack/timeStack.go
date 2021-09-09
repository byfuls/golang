package main

import (
	"fmt"
)

func timeStackInit(size int) *timeStack {
	return &timeStack{
		stack: make([]uint, size),
		size:  size,
	}
}

type timeStack struct {
	stack []uint
	size  int
	head  int
	tail  int
	count int
}

func (q *timeStack) Push(n uint) {
	if q.head == q.tail && q.count > 0 {
		stack := make([]uint, len(q.stack)+q.size)
		copy(stack, q.stack[q.head:])
		copy(stack[len(q.stack)-q.head:], q.stack[:q.head])
		q.head = 0
		q.tail = len(q.stack)
		q.stack = stack
	}
	q.stack[q.tail] = n
	q.tail = (q.tail + 1) % len(q.stack)
	q.count++
}

func (q *timeStack) Pop() uint {
	if q.count == 0 {
		return 0
	}
	stack := q.stack[q.head]
	q.head = (q.head + 1) % len(q.stack)
	q.count--
	return stack
}

func main() {
	q := timeStackInit(1)
	q.Push(4)
	q.Push(5)
	q.Push(6)
	fmt.Println(q.Pop(), q.Pop(), q.Pop())
}

