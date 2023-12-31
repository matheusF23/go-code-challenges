package main

import "fmt"

func isValid(s string) bool {
	charMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []rune{}
	for _, char := range s {
		if _, ok := charMap[char]; ok {
			stack = append(stack, char)
		} else {
			if len(stack) > 0 && char == charMap[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	s := "([]}])"
	fmt.Println(isValid(s))
}
