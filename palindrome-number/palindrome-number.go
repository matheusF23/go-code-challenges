package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	var xString = strconv.Itoa(x)
	var i = 0
	var j = len(xString) - 1
	for i < j {
		if xString[i] != xString[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(1212))
}
