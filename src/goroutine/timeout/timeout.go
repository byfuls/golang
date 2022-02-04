// ING ING ING ING ING ING ING ING ING ING
package main

import (
	"context"
	"fmt"
	"time"
)

type Response struct {
	data   interface{}
	status bool
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))

	func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
		default:
		}

		time.Sleep(1 * time.Second)
	}(ctx)
	defer cancel()

	for {
		time.Sleep(1 * time.Second)
	}
}
