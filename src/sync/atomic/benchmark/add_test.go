package add

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	var ops uint64 = 0
	var wg sync.WaitGroup

	// fmt.Println("start")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < b.N; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	_ = atomic.LoadUint64(&ops)
	// opsFinal := atomic.LoadUint64(&ops)
	// fmt.Println("done, ops: ", opsFinal)
}
