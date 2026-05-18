package main

import (
	"fmt"
	"time"
)
	
func numOne() {
	for i := 1; i <= 5; i++ {
		fmt.Println("one:", i)
	}
}
func numTwo() {
	for i := 1; i <= 5; i++ {
		fmt.Println("two:", i)
	}
}
func numThree() {
	for i := 1; i <= 5; i++ {
		fmt.Println("three:", i)
	}
}

func main() {
	go numOne()
	go numTwo()
	go numThree()
	time.Sleep(time.Second)
}



