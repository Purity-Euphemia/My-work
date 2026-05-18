package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const height = 8

// 🎨 ANSI colors
func getColor(c string) string {
	switch strings.ToLower(c) {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	default:
		return "\033[0m"
	}
}

// 📂 Load banner file
func loadBanner(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

// 🔤 Get ASCII lines for a character safely
func getCharArt(banner []string, r rune) []string {
	if r < 32 || r > 126 {
		return make([]string, height)
	}
	start := (int(r) - 32) * height

	if start+height > len(banner) {
		return make([]string, height)
	}

	return banner[start : start+height]
}

// 🎯 Check if substring starts at position
func matchAt(text, sub string, i int) bool {
	return i+len(sub) <= len(text) && text[i:i+len(sub)] == sub
}

func main() {

	// ✅ Validate arguments
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")
		return
	}

	args := os.Args[1:]

	var color, sub, text string

	// 🎯 Parse arguments
	if strings.HasPrefix(args[0], "--color=") {
		color = strings.TrimPrefix(args[0], "--color=")

		if len(args) == 3 {
			sub = args[1]
			text = args[2]
		} else if len(args) == 2 {
			text = args[1]
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			return
		}
	} else {
		text = args[0]
	}

	// 🧹 Clean input
	text = strings.TrimSpace(text)
	sub = strings.TrimSpace(sub)

	// 📂 Load banner
	banner, err := loadBanner("banners/standard.txt")
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}

	colorCode := getColor(color)
	reset := "\033[0m"

	// 🔁 Print ASCII art
	for line := 0; line < height; line++ {

		i := 0
		for i < len(text) {

			// 🎨 Color substring
			if sub != "" && matchAt(text, sub, i) {

				for j := 0; j < len(sub); j++ {
					charArt := getCharArt(banner, rune(text[i+j]))
					fmt.Print(colorCode + charArt[line] + reset)
				}
				i += len(sub)

			} else {

				charArt := getCharArt(banner, rune(text[i]))

				// 🎨 Color whole text if no substring
				if sub == "" && color != "" {
					fmt.Print(colorCode + charArt[line] + reset)
				} else {
					fmt.Print(charArt[line])
				}

				i++
			}
		}
		fmt.Println()
	}
}