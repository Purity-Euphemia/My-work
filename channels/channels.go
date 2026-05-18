package main

import "fmt"

func channel() {
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	value := <-ch
	fmt.Println(value)
}
