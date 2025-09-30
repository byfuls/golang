package main

import (
	"fmt"
	"strings"
)

const (
	Y = 'Y' // Yellow jelly
	P = 'P' // Purple jelly
	I = 'I' // Ice jelly
)

// jellyTransformation solves the jelly transformation problem
// bears: string representing jellies ('Y', 'P', 'I')
// K: number of consecutive jellies that can be transformed at once
// Returns minimum number of reagents needed, or -1 if impossible
func jellyTransformation(bears string, K int) int {

	if K <= 0 || K > len(bears) {
		return -1
	}

	// Check if already all ice jellies
	if allIce(bears) {
		return 0
	}

	// BFS to find minimum steps
	queue := []string{bears}
	visited := make(map[string]bool)
	visited[bears] = true
	steps := make(map[string]int)
	steps[bears] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentSteps := steps[current]

		// Try all possible K-length consecutive segments
		for i := 0; i <= len(current)-K; i++ {
			next := transformSegment(current, i, K)

			if !visited[next] {
				visited[next] = true
				steps[next] = currentSteps + 1

				// Check if we reached the goal
				if allIce(next) {
					return steps[next]
				}

				queue = append(queue, next)
			}
		}
	}

	// If we can't reach all ice state
	return -1
}

// transformSegment transforms K consecutive jellies starting from index start
func transformSegment(bears string, start, K int) string {
	result := []rune(bears)

	for i := start; i < start+K; i++ {
		switch result[i] {
		case 'Y':
			result[i] = P
		case 'P':
			result[i] = I
		case 'I':
			result[i] = Y
		}
	}

	return string(result)
}

// allIce checks if all jellies are ice jellies
func allIce(bears string) bool {
	return !strings.Contains(bears, "Y") && !strings.Contains(bears, "P")
}

func main() {
	// Test cases
	testCases := []struct {
		bears string
		K     int
		want  int
	}{
		{"IPYIYP", 3, 3},
		{"IY", 1, 2},
		{"PPY", 2, -1},
	}

	for i, tc := range testCases {
		result := jellyTransformation(tc.bears, tc.K)
		fmt.Printf("Test case %d: bears=%s, K=%d, result=%d, expected=%d\n",
			i+1, tc.bears, tc.K, result, tc.want)
	}
}
