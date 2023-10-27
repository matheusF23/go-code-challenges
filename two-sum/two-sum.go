package main

import "fmt"

// CHALLENGE: https://leetcode.com/problems/two-sum/

// func twoSum(nums []int, target int) []int {
// 	var result []int
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				result = []int{i, j}
// 				return result
// 			}
// 		}
// 	}
// 	return result
// }

func twoSum(nums []int, target int) []int {
	mapper := map[int]int{}
	for j, num := range nums {
		complementary := target - num
		if i, ok := mapper[num]; ok {
			return []int{i, j}
		}
		mapper[complementary] = mapper[complementary] | j
	}
	return []int{}
}

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}
