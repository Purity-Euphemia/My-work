package main

//import "fmt"

func palindrome(word string) bool {
	left := 0
	right := len(word) - 1

	for left < right{
		if word[left] != word[right] {
			return false
		}
		left++
		right++
	}
	return true
}
