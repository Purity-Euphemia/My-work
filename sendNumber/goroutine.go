package main

import (
	"fmt"
)

func sendNumber(ch chan int) {
	ch <- 10
}

func main() {
	ch := make(chan int)

	go sendNumber(ch)

	value := <-ch
	fmt.Println(value)
}
