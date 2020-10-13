package main

import "fmt"

/* QUEUE (+slice)
put new data, O(N) (capacity가 크다면 다르겠지만 cap과 len이 같다면,)
get data, O(1)
*/

func main() {
	queue := []int{}
	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}

	fmt.Println(queue)

	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}
}
