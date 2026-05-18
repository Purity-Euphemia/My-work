package main

import "fmt"

func factorial(s int) int { 
	a := 1
	for i := 1; i <= s; i++ {
		a *= i
	}
	return a
}

func main() {
	fmt.Println(factorial(5))
}
