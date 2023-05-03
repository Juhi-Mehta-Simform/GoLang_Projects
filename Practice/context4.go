package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	doSomethhing(ctx)
}

func doSomethhing(ctx context.Context) {
	ctx, cancelCtx := context.WithTimeout(ctx, 1500*time.Millisecond)
	printCh := make(chan int)
	go doAnothher(ctx, printCh)
	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}
	defer cancelCtx()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("doSomething -> finished")
}

func doAnothher(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Println("doAnother Finished")
			return
		case num := <-printCh:
			fmt.Printf("doAnother : %d\n", num)
		}
	}
}
