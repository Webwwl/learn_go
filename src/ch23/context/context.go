package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	receive(ctx)
	cancel()
	time.Sleep(time.Second * 1)
}

func receive(ctx context.Context) {
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println("canceled:", i)
		}(i, ctx)
	}
}

func isCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
