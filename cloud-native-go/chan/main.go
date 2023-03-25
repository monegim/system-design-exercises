package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println("Sending value to channel complete")
	val := <-ch
	val = <- ch
	fmt.Printf("Value received: %d\n", val)
}
