package main

import "fmt"

func main() {
	x := 5
	increment := func() int {
		x = x + 1
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
