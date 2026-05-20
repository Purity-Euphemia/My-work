package main

import "fmt"

func removeDuplicate(nums []int) []int {
	value := []int{}

	for count := 0; count < len(nums); count++{
		found := false

		for counter := 0; counter < len(value); counter++{
			if nums[count] == value[counter] {
				found = true
			}
		}
		if !found {
			value = append(value, nums[count])
		}
	}
	return value
}

func main() {
	numbers := []int{1, 2, 2, 4, 4, 3}
	fmt.Println(removeDuplicate(numbers))
}
