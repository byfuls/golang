package main

import "fmt"

/* stack (+slice)
put new data, O(N) (capacity가 크다면 다르겠지만 cap과 len이 같다면,)
get data, O(1)
*/

func main() {
	stack := []int{}

	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}

	fmt.Println(stack)

	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}
}
