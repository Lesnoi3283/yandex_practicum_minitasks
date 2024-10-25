package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

//task: Задача в том, чтобы предусмотреть механизм аккуратного завершения работы при прерывании. Допишите код.

func main() {
	//code from task:
	var srv = http.Server{Addr: ":8080"}
	// сервер минимальный
	// не будем даже декларировать и регистрировать обработчики
	// для задачи они не важны

	//my code:
	closeCh := make(chan os.Signal, 1)
	signal.Notify(closeCh, os.Interrupt)

	wg := sync.WaitGroup{}
	wg.Add(1)
	//os signals listen func:
	go func() {
		defer wg.Done()
	loop:
		for {
			select {
			case <-closeCh:
				log.Printf("Shutting down")
				err := srv.Shutdown(context.Background())
				if err != nil {
					log.Printf("HTTP server Shutdown: %v", err)
				}
				break loop

			default:
				//nothing
			}
		}
	}()

	//code from task:
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// ошибки запуска или остановки Listener
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	//my code:
	wg.Wait()
	close(closeCh)
}
