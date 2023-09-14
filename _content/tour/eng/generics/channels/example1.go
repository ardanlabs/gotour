//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to execute a work function in a goroutine and
// return a channel of type Result (to be determined later) back to the caller.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type doworkFn[Result any] func(context.Context) Result

func doWork[Result any](ctx context.Context, work doworkFn[Result]) chan Result {
	ch := make(chan Result, 1)

	go func() {
		ch <- work(ctx)
		fmt.Println("doWork : work complete")
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	dwf := func(ctx context.Context) string {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		return "work complete"
	}

	select {
	case v := <-doWork(ctx, dwf):
		fmt.Println("main:", v)
	case <-ctx.Done():
		fmt.Println("main: timeout")
	}
}
