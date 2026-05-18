package main

import "fmt"

func getLarge(nums []int) (int, int){
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
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		large, secondLargest := getLarge(nums)

		fmt.Println("Largest:", large)
		fmt.Println("Second Largest:", secondLargest)
	}
