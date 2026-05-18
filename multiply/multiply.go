package main

import "fmt"

func multiply(n int, s int) int {
	value := n * s
	return value
}

func main() {
	s := 3
	t := 6
	fmt.Println(multiply(s, t))
}
