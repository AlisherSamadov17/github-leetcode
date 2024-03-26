package main

import "fmt"

func lengthOfLastWord(s string) int {
	l := 0
	sum := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			l++
			sum = true
		} else if sum {
			break
		}
	}
	return l
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}
