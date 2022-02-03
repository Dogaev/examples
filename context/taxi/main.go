/*
	Пример context с отменой.
	Одновременно вызываем такси в нескольких сервисах, нужно вывести
	сервис такси, где машина нашлась быстрее и остановить поиск в
	остальных сервисах
*/
package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var (
		resultCh     = make(chan string)
		ctx, cancel  = context.WithCancel(context.Background())
		taxiServices = []string{"uber", "yandexGO", "ситимобайл", "redtaxi"}
		wg           sync.WaitGroup
		winner       string
	)

	defer cancel()

	for i := range taxiServices {
		svc := taxiServices[i]

		wg.Add(1)
		go func() {
			requestRide(ctx, svc, resultCh)
			wg.Done()
		}()
	}

	go func() {
		winner = <-resultCh
		cancel() // преднамеренно отменяем контекст, чтобы остальные горутины получили сигнал, что нужно прекратить поиск машины
	}()

	wg.Wait()
	log.Printf("found car in %q", winner)
}

func requestRide(ctx context.Context, serviceName string, resultCh chan string) {

	for {
		select {
		case <-ctx.Done():
			log.Printf("stopped the search in %q (%v)", serviceName, ctx.Err())
			return
		default:
			if rand.Float64() > 0.90 {
				resultCh <- serviceName
				return
			}

			time.Sleep(1*time.Millisecond)
			continue
		}
	}
}
