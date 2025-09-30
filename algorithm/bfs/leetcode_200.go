package bfs

func numIslands(grid [][]byte) int {
	var (
		rows, cols = len(grid), len(grid[0])
		queue      [][]int
		dirs       = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		result     int
	)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				result++
				grid[r][c] = '0' // Mark as visited
				queue = append(queue, []int{r, c})
				for len(queue) > 0 {
					point := queue[0]
					queue = queue[1:]
					_r, _c := point[0], point[1]
					for _, d := range dirs {
						_mr, _mc := _r+d[0], _c+d[1]
						if 0 > _mr || _mr >= rows ||
							0 > _mc || _mc >= cols || grid[_mr][_mc] == '0' {
							continue
						}

						grid[_mr][_mc] = '0' // Mark as visited
						queue = append(queue, []int{_mr, _mc})
					}
				}
			}
		}
	}
	return result
}
