package main

func capitalize(word string) string{
	result := ""

	for count := 0; count < len(word); count++{
		letter := word[count]

		if letter >= 'a' && letter <= 'z' {
			letter = letter - 32
		}
		result += string(letter)
	}
	return result
}
