package dfs

func numIslands(grid [][]byte) int {
	var (
		dfs        func(r, c int)
		rows, cols = len(grid), len(grid[0])
		result     int
	)

	dfs = func(r, c int) {
		if 0 > r || r >= rows ||
			0 > c || c >= cols || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'
		dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, d := range dirs {
			dfs(r+d[0], c+d[1])
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				result++
				dfs(r, c)
			}
		}
	}

	return result
}
