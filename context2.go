package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	doSomethingg(ctx)
}

func doSomethingg(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)
	printCh := make(chan int)
	go doAnotheer(ctx, printCh)
	for num := 1; num <= 3; num++ {
		printCh <- num
	}
	cancelCtx()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("DoSomething -> finished")
}

func doAnotheer(ctx context.Context, printCh <-chan int) {
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
