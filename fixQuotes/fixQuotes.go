package main

import "regexp"


func fixQuotes(text string) string {
	re := regexp.MustCompile(`'\s+(.+?)\s+'`)
	return re.ReplaceAllString(text, "'$1'")
}
