package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func in(ctx context.Context, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var a string
	for {
		select {
		case <-ctx.Done():
			log.Println("Context cancelled, stop reading")
			close(ch)
			return
		default:
			fmt.Scan(&a)
			ch <- a
		}
	}

}

func out(ctx context.Context, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	f, _ := os.Create("file.txt")
	defer f.Close()

	for {
		select {
		case <-ctx.Done():
			for v := range ch {
				f.WriteString(v)
			}
			log.Println("Context cancelled, stopped writing")
			return
		}

	}
}

func main() {
	var (
		resultCh    = make(chan string)
		ctx, cancel = context.WithCancel(context.Background())
		wg          sync.WaitGroup
	)

	wg.Add(2)
	go in(ctx, resultCh, &wg)
	go out(ctx, resultCh, &wg)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	go func() {
		<-sigChan
		cancel()
	}()

	wg.Wait()
}
