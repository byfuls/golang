package child

import "testing"

func TestInit(t *testing.T) {
	p1 := Init()
	p1.Debug()

	p2 := Init()
	p2.Debug()
}
