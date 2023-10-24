package main

import (
	"fmt"
)

var dict = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	soma := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && dict[string(s[i])] < dict[string(s[i+1])] {
			soma -= dict[string(s[i])]
		} else {
			soma += dict[string(s[i])]
		}
	}
	return soma
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
