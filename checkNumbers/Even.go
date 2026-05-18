package main

import "fmt"

func CheckNumbers(g int) string {
	if g % 2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	value := 7
	fmt.Println(CheckNumbers(value))
}
