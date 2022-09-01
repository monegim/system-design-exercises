package main

import (
	"fmt"
	"time"
)

func main(){
	t := time.Now()
	t_after := t.Add(time.Second * 2)
	a := 2
	fmt.Println("Now", t)
	fmt.Printf("After %d: %s\n",a, t_after)
	fmt.Println(time.Now().After(t_after))
}