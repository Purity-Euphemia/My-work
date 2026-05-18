package main

import (
	"fmt"
	"time"
)


func printNumber() {
	for i := 1; i <= 5; i++{
		fmt.Println(i)
	}
}

func main() {
	go printNumber()
	time.Sleep(time.Second)
}