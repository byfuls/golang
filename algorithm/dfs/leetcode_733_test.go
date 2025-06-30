package dfs

import (
	"testing"
)

func TestQ733floodFill(t *testing.T) {
	var (
		image    = [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
		sr       = 1
		sc       = 1
		newColor = 2
		result   [][]int
	)
	result = Q733floodFill(image, sr, sc, newColor)
	t.Log(result)
}
