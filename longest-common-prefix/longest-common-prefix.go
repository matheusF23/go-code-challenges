package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	commonPrefixes := []byte{}
	for i := 0; i < 200; i++ {
		if i == len(strs[0]) {
			break
		}
		commonPrefixes = append(commonPrefixes, strs[0][i])
		for _, str := range strs[1:] {
			if i == len(str) {
				return string(commonPrefixes[:len(commonPrefixes)-1])
			}
			if str[i] != commonPrefixes[len(commonPrefixes)-1] {
				return string(commonPrefixes[:len(commonPrefixes)-1])
			}
		}
	}
	return string(commonPrefixes)
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
