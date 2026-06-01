package main

import "fmt"

func main() {
	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))
}
