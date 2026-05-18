package main

import "fmt"

func main() {
	fmt.Println(IfElse(5))
}

func IfElse(even int) bool {
		if even % 2 == 0 {
			return true
	}
	return false
}
