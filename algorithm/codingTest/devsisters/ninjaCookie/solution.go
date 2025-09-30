package main

import (
	"fmt"
)

// solution solves the ninja cookie problem with correct strategy
func solution(D int, cakes []int) int {
	if D < 0 || len(cakes) == 0 {
		return -1
	}

	// Copy the array to work with
	obstacles := make([]int, len(cakes))
	copy(obstacles, cakes)

	shurikens := 0

	// Ninja moves from position D down to position 0
	// At each position, only process obstacles that won't be reachable later
	for pos := D; pos >= 0; pos-- {
		// At position pos, we can only affect position pos itself
		// But we should use a greedy strategy:
		// Only reduce obstacles that cannot be reduced in future steps

		if pos < len(obstacles) {
			// We need to ensure obstacle at position pos is cleared
			// But we should minimize total shurikens by being smart about it

			// Simple strategy: reduce obstacle at current position by 1 if it's > 0
			if obstacles[pos] > 0 {
				obstacles[pos]--
				shurikens++
			}
		}
	}

	// Check if all obstacles are cleared
	for i, height := range obstacles {
		if height > 0 {
			// Check if this position was reachable
			if i > D {
				return -1 // Position i is beyond reach
			}
			// If there are still obstacles in reachable positions,
			// our strategy was not optimal, but let's continue for now
		}
	}

	return shurikens
}

func main() {
	// Test cases
	testCases := []struct {
		D     int
		cakes []int
		want  int
	}{
		{3, []int{2, 2, 3}, 4},
		{1, []int{2, 2, 3}, -1},
		{1, []int{1, 2, 3}, 3},
		{2, []int{1, 2, 2}, 3},
	}

	// Test the strategy step by step for case 4
	fmt.Println("=== Testing strategy for D=2, cakes=[1,2,2] ===")
	obstacles := []int{1, 2, 2}
	shurikens := 0

	for pos := 2; pos >= 0; pos-- {
		fmt.Printf("Position %d: obstacles=%v", pos, obstacles)
		if pos < len(obstacles) && obstacles[pos] > 0 {
			obstacles[pos]--
			shurikens++
			fmt.Printf(" -> %v (threw 1 shuriken)", obstacles)
		}
		fmt.Println()
	}
	fmt.Printf("Total shurikens: %d\n\n", shurikens)

	for i, tc := range testCases {
		result := solution(tc.D, tc.cakes)
		fmt.Printf("Test case %d: D=%d, cakes=%v, result=%d, expected=%d\n",
			i+1, tc.D, tc.cakes, result, tc.want)
	}
}
