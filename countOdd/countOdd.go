package main

import "fmt"

func countOdd(arr []int) int {
	value := 0
	for count := 0; count < len(arr); count++{
		if arr[count] % 2 != 0{
			value++
		}
	}
	return value
}

func main() {
	number := []int{3, 8, 11, 14, 7, 2}
	fmt.Println(countOdd(number))
}
