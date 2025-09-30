package bfs

import "testing"

func TestFloodFill(t *testing.T) {
	var (
		image    = [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
		sr       = 1
		sc       = 1
		newColor = 2
		result   [][]int
	)
	result = floodFill(image, sr, sc, newColor)
	t.Log(result)
}
