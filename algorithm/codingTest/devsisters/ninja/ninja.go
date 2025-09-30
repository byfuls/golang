package main

import (
	"fmt"
)

func solution(D int, cakes []int) int {
	n := len(cakes)
	throws := make([]int, n) // 각 위치에서 던진 표창 수
	curr := 0                // 현재까지 누적된 표창 영향
	result := 0

	for i := 0; i < n; i++ {
		// D칸을 벗어난 과거 영향 제거
		if i >= D {
			curr -= throws[i-D]
		}

		// 현재 위치에서 필요한 만큼 추가로 던짐
		remain := cakes[i] - curr
		if remain > 0 {
			if i+D > n {
				return -1
			}
			throws[i] = remain
			curr += remain
			result += remain
		}
	}

	return result
}

func main() {
	fmt.Println(solution(3, []int{2, 2, 3})) // ✅ 4
	fmt.Println(solution(1, []int{2, 2, 3})) // ✅ -1
	fmt.Println(solution(2, []int{2, 2, 3})) // ✅ 3
	fmt.Println(solution(1, []int{1, 2, 3})) // ✅ -1
	fmt.Println(solution(2, []int{1, 2, 2})) // ✅ 3
}
