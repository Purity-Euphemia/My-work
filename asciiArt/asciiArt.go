package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	lines := strings.Split(input, "\\n")

	banner := loadBanner("standard.txt")

	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		printLine(line, banner)
	}
}

func loadBanner(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil
	}
	return strings.Split(string(data), "\n")
}

func printLine(text string, banner []string) {
	for row := 0; row < 8; row++{
		for _, ch := range text {
			if ch < 32 || ch > 126 {
				continue
			}
			index := (int(ch) - 32) * 9
			fmt.Print(banner[index+row])
		}
		fmt.Println()
	}
}