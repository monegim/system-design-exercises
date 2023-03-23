package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go send(ch)

	go receive(ch)
	time.Sleep(2 * time.Second)
}

func send(ch chan int) {
	ch <- 1
	fmt.Println("Sending value to channel complete")
}

func receive(ch chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Timeout finished")

	_ = <- ch
	return
}
