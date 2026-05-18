package main

import "fmt"

func palindrome(s string) bool {
	for i := 1; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	c := "verse"
	fmt.Println(palindrome(c))
}
