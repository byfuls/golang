package main

import (
	"fmt"
)

// jellyTransformation solves the jelly transformation problem with optimized approach
// bears: string representing jellies ('Y', 'P', 'I')
// K: number of consecutive jellies that can be transformed at once
// Returns minimum number of reagents needed, or -1 if impossible
func jellyTransformation(bears string, K int) int {
	// Check if already all ice jellies (including empty string)
	if allIce(bears) {
		return 0
	}

	if K <= 0 || K > len(bears) {
		return -1
	}

	// For K=1, we can use mathematical approach
	if K == 1 {
		return calculateStepsForK1(bears)
	}

	// For K > 1, use optimized BFS
	return bfsOptimized(bears, K)
}

// calculateStepsForK1 calculates minimum steps when K=1
func calculateStepsForK1(bears string) int {
	totalSteps := 0

	for _, char := range bears {
		switch char {
		case 'Y':
			totalSteps += 2 // Y -> P -> I
		case 'P':
			totalSteps += 1 // P -> I
		case 'I':
			totalSteps += 0 // Already I
		}
	}

	return totalSteps
}

// bfsOptimized performs optimized BFS for K > 1
func bfsOptimized(bears string, K int) int {
	queue := []string{bears}
	visited := make(map[string]bool)
	visited[bears] = true
	steps := make(map[string]int)
	steps[bears] = 0

	// Limit iterations to prevent timeout
	maxIterations := 5000
	iterations := 0

	for len(queue) > 0 && iterations < maxIterations {
		current := queue[0]
		queue = queue[1:]
		currentSteps := steps[current]
		iterations++

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

				// Limit queue size to prevent memory issues
				if len(queue) < 500 {
					queue = append(queue, next)
				}
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
			result[i] = 'P'
		case 'P':
			result[i] = 'I'
		case 'I':
			result[i] = 'Y'
		}
	}

	return string(result)
}

// allIce checks if all jellies are ice jellies
func allIce(bears string) bool {
	// Empty string is considered all ice
	if len(bears) == 0 {
		return true
	}

	for _, char := range bears {
		if char != 'I' {
			return false
		}
	}
	return true
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
