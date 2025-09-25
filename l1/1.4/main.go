package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/lastlife77/wb-tech/l1/1.4/handler"
)

func main() {
	var wg sync.WaitGroup

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	h := handler.New(getStreams())
	h.HandleMsgs(ctx, &wg)

	<-ctx.Done()
	fmt.Println("||||||||||||||||||||||")
	fmt.Println("Идет завершение работы")
	fmt.Println("||||||||||||||||||||||")
	wg.Wait()
}

func getStreams() int {
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: не задано количество потоков")
		fmt.Println("Использование: program <количество_потоков>")
		os.Exit(1)
	}

	streams, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Ошибка: аргумент '%s' не является числом \n", os.Args[1])
		os.Exit(1)
	}
	return streams
}
