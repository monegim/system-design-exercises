package main

import (
	"context"
	"fmt"
)

func main() {
	new_context := context.Background()
	fmt.Printf("%T",new_context)
}