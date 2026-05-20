package main

import "fmt"

func secondLargest(nums []int) (int, int) {
	large := nums[0]
	secondLargest := nums[0]

	for count := 0; count < len(nums); count++ {
		if nums[count] > large {
			secondLargest = large
			large = nums[count]
		} else if nums[count] > secondLargest && nums[count] != large {
			secondLargest = nums[count]
		}
	}
	return large, secondLargest
}

func main() {
	numbers := [] int{5, 2, 10, 4, 20}
	fmt.Println(secondLargest(numbers))
}


