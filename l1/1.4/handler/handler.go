package handler

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Handler struct {
	msgCount int
	streams  int
}

func New(streams int) *Handler {
	return &Handler{
		msgCount: 0,
		streams:  streams,
	}
}

func (h *Handler) HandleMsgs(ctx context.Context, wg *sync.WaitGroup) {
	ch := make(chan string, 100)

	wg.Add(1)
	go h.writeMsg(ctx, ch, wg)
	for i := 1; i <= h.streams; i++ {
		wg.Add(1)
		go h.readMsg(ch, i, wg)
	}

}

func (h *Handler) writeMsg(ctx context.Context, ch chan<- string, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case ch <- randomMsg():
			time.Sleep(1 * time.Second)
			h.msgCount++
			fmt.Print("A message was written. ")
			fmt.Printf("Msgs count: %d \n", h.msgCount)
		}
	}

}

func (h *Handler) readMsg(ch <-chan string, index int, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range ch {
		time.Sleep(5 * time.Second)
		h.msgCount--
		fmt.Printf("Worker %d: A message was read: %s. ", index, msg)
		fmt.Printf("Msgs count: %d \n", h.msgCount)
	}
}

func randomMsg() string {
	return strconv.Itoa(rand.Intn(100))
}
