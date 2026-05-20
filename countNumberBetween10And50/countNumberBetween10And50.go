package main

import "fmt"

func countNumber(arr []int) int{
	value := 0
	for count := 0; count < len(arr); count++{
		if arr[count] >= 10 && arr[count] <= 50 {
			value++
		}
	}
	return value
}

func main() {
	number := []int{5, 12, 45, 60, 9, 30, 35}
	fmt.Println(countNumber(number))
}
