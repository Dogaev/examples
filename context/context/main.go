package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	doWork(ctx)
}

// функция doWork будет отрабатывать кол-во сек которые указаны в
// родительском контексте, тот таймаут, что указан в контексте данной
// функции (newCtx) не будет учтен в данной програме, так как родительский контекст
// завершится быстрее
func doWork(ctx context.Context) {
	newCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	log.Println("starting working....")

	for {
		select {
		case <-newCtx.Done():
			log.Printf("ctx done: %v", ctx.Err())
			return
		default:
			log.Println("working...")
			time.Sleep(time.Second * 1)
		}
	}
}
