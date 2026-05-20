package main

import "fmt"

func countNegative(arr []int) int {
	value := 0
	for count := 0; count < len(arr); count++{
		if arr[count] < 0 {
			value++
		}
	}
	return value
}

func main() {
	number := []int{4, -1, 6, -7, 0, -3, 9}
	fmt.Println(countNegative(number))
}
