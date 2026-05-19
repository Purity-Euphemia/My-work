package main

import "fmt"

func smallestNumber(nums []int) int {
	smallestNumber := nums[0]
	for count := 0; count < len(nums); count++{
		if nums[count] < smallestNumber {
			smallestNumber = nums[count]
		}
	}
	return smallestNumber
}

func main() {
	numbers := [] int{20, 30, 40, 50}
	fmt.Println(smallestNumber(numbers))
}
