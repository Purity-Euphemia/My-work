package main

import (
	"fmt"
	"strings"
)

// func fixAtoAn(text string) string{
// 	words := strings.Fields(text)

// 	for count := 0; count < len(words)-1; count++{
// 		word := words[count]
// 		next := words[count+1]

// 		if (word == "a" || word == "A")
// 	}


func fixAtoAn(s string)string{
	word := ""
	for count := 0; count < len(s); count++{
		if s[count] == 'a' && count + 2 < len(s) && s[count + 1] == ' ' && strings.ContainsRune("aeiou",rune(s[count])){
			word += "an"
		} else if s[count] == 'A' && count + 2 < len(s) && s[count + 1] == ' ' && strings.ContainsRune("AEIOU", rune(s[count])){
			word += "An"
		}else {
			word += string(s[count])
		}
	}
	return word
}

func main() {
	number := "i am a apple is good"
	fmt.Println(fixAtoAn(number))
}
