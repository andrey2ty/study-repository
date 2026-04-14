package pid

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func SignInt() {
	pid := os.Getpid()
	fmt.Println("pid:", pid)

	sign := make(chan os.Signal, 1)
	term := make(chan os.Signal, 1)

	signal.Notify(sign, syscall.SIGINT)
	signal.Notify(term, syscall.SIGTERM)

	go func() {
		for sig := range sign {
			fmt.Println("отловлен сигнал", sig)
		}
	}()

	<-term
	fmt.Println("Завершение")
}

func WorkerEnv() {
	pid := os.Getpid()
	fmt.Println("pid:", pid)
	str := os.Getenv("WORKERS_COUNT")

	count, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	for i := 1; i <= count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Println("I'm worker number", i)
					time.Sleep(time.Second)
				}
			}
		}()
	}

	sInt := make(chan os.Signal, 1)
	sTerm := make(chan os.Signal, 1)

	signal.Notify(sInt, syscall.SIGINT)
	signal.Notify(sTerm, syscall.SIGTERM)

	go func() {
		for {
			select {
			case sign := <-sInt:
				fmt.Println("отловлен", sign)
			case <-sTerm:
				fmt.Println("Был вызван SIGTERM, остановка...")
				cancel()
			}
		}
	}()

	wg.Wait()
	fmt.Println("все остановилось")
}
