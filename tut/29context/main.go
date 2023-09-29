package main

import (
	"context"
	"fmt"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

func doSomething(ctx context.Context) {
	rID := ctx.Value("request-id")
	fmt.Println(rID)
}

func main() {
	fmt.Println("Context in GO")
	ctx := context.Background()
	ctx = enrichContext(ctx)
	doSomething(ctx)
}
