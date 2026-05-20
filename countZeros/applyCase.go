package main

import (
	"strings"
	"unicode"
)

func applyCase(word, cmd string) string{
	switch cmd {
	case "up":
		return  strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	case "cap":
		if len(word) == 0 {
			return word
		}
		r := []rune(word)
		return string(unicode.ToUpper(r[0]))
		strings.ToLower(string(r[1:]))
	}
	return word
}