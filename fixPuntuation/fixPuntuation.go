package main

import (
	"regexp"
	"strings"
)

func fixPuntuation(text string) string {
	rep1 := regexp.MustCompile(` +([.,!?:;]+)`)
	text = rep1.ReplaceAllString(text, "$1")

	rep2 := regexp.MustCompile(` +([.,!?:;]+)`)
	text = rep2.ReplaceAllString(text, "$1 $2")

	return strings.TrimSpace(text)
}
