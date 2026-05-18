package main

import (
	"fmt"
	"sync"
)

func printNumbers(name string, wg *sync.WaitGroup){
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(name, i)
	}
}
func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go printNumbers("one:", &wg)
	go printNumbers("two:", &wg)
	go printNumbers("three:", &wg)

	wg.Wait()
}
