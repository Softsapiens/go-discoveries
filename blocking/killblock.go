package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

func main() {

	ctx := context.Background()
	cancelctx, cancel := context.WithCancel(ctx)

	go blocker(cancelctx)

	fmt.Println("Press enter to cancel")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')

	fmt.Println("Cancelling work...")
	cancel()

	fmt.Println("Program End")
}

func blocker(ctx context.Context) {
	fmt.Println("for block started")

	done := false
	for i := 0; i < 10000000000 && !done; i++ {

		if i % 1000000 == 0 {
			select {
			case <-ctx.Done():
				done = true
			default:
			}
		}
	}
	fmt.Println("for block finished")
}
