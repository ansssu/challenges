package main

import (
	"fmt"
	"os"
)

func IsBalanced(text string) bool {

	s := make([]rune, 0, len(text))

	for _, c := range text {
		if c == '(' {
			s = append(s, c)
		} else if c == ')' {
			if len(s) == 0 {
				return false
			}
			s = s[:len(s)-1]
		}

		fmt.Println(string(s))
	}

	if len(s) != 0 {
		return true
	}

	return false
}

func main() {

	input := os.Args[1]
	var result string
	if IsBalanced(input) {
		result = "balanced"
	} else {
		result = "not balanced"
	}

	fmt.Println(result)

}
