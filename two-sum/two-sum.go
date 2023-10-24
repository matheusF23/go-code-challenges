package main

import "fmt"

// CHALLENGE: https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
				return result
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 4, 6, 2, 5}
	target := 7
	fmt.Println(twoSum(nums, target))
}
