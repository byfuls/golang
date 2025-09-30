package bfs

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	originColor := image[sr][sc]
	if originColor == newColor {
		return image
	}

	rows, cols := len(image), len(image[0])
	queue := [][]int{{sr, sc}}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		r, c := point[0], point[1]

		image[r][c] = newColor

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows &&
				nc >= 0 && nc < cols &&
				image[nr][nc] == originColor {
				queue = append(queue, []int{nr, nc})
			}
		}
	}

	return image
}
