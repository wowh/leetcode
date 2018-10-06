package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")
	if len(words) > 0 {
		for i := len(words) - 1; i >= 0; i-- {
			if len(words[i]) == 0 {
				continue
			}
			return len(words[i])
		}
	}

	return 0
}

func main() {
	fmt.Println(lengthOfLastWord("a "))
}
