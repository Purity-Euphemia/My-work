package main

import "fmt"

func largestNumber(arr []int) int {
	largestNumber := arr[0]
	for count := 0; count < len(arr); count++{
		if arr[count] > largestNumber {
			largestNumber = arr[count]
		}
	}
	return largestNumber
}

func main() {
	number := [] int{10, 20, 30, 40, 50}
	fmt.Println(largestNumber(number))
}
