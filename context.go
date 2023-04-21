package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Println("Doing Something")
	fmt.Printf("DoSomething: PIN's Value is %s\n", ctx.Value("PIN"))
	anotherCtx := context.WithValue(ctx, "PIN", "0000")
	doAnother(anotherCtx)
}

func doAnother(ctx context.Context) {
	fmt.Printf("DoAnother: Pin is %s\n", ctx.Value("PIN"))
}

func main() {
	// ctx := context.TODO()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "PIN", "6402")
	doSomething(ctx)
}
