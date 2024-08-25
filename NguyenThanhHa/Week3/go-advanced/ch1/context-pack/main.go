package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) error{
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <- ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	// Nhan vao context parent (backbround) va tra ve context child (ctx) va ham cancel
	// deadline 10 secs
	ctx, cancel := context.WithTimeout(context.Background(),10 * time.Second)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	time.Sleep(time.Second)

	cancel()

	wg.Wait()
}