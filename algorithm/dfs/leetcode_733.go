package dfs

// Leetcode 733 - Flood Fill
// URL : https://leetcode.com/problems/flood-fill/description/

func Q733floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var originColor = image[sr][sc]
	if originColor == newColor {
		return image
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if 0 > r || r >= len(image) ||
			0 > c || c >= len(image[0]) {
			return
		}

		if image[r][c] != originColor {
			return
		}

		image[r][c] = newColor
		dfs(r-1, c) // up
		dfs(r+1, c) // down
		dfs(r, c-1) // left
		dfs(r, c+1) // right
	}

	dfs(sr, sc)
	return image
}
