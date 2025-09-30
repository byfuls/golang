package dfs

import "testing"

func TestNumIslands(t *testing.T) {
	var (
		grid = [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}
		result int
	)
	result = numIslands(grid)
	t.Log(result) // Expected output: 1
}
