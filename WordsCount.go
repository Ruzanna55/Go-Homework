package main

import "fmt"

func main() {

	fmt.Println("This Program returns word count in the given string!")
	var sentence string = "I Love Programming."
	c := wordsCount(sentence)
	fmt.Println("The number  of words is", c)
}
func wordsCount(sentence string) int {
	fmt.Println("The sentence is:")
	fmt.Println(sentence)
	x := 0
	count := 0
	for x < len(sentence) {
		i := string([]rune(sentence)[x])
		if i == " " || i == "." || i == "?" || i == "!" {
			count++
		}

		x++
	}

	return count
}
