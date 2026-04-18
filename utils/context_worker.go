package utils

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d, ruk gaya \n", id)
			return

		default:
			fmt.Printf("Workern %d kaam kar raha hai \n", id)
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func StartWorker() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, 1)
	go worker(ctx, 2)
	go worker(ctx, 3)

	time.Sleep(time.Second * 4)
	fmt.Println("Boss - sb kamm band karo")
	cancel()

	time.Sleep(1 * time.Second)
}
