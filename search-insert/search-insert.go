package main

import "fmt"

// O(log n) solution
func searchInsert(nums []int, target int) int {
	base := 0
	top := len(nums) - 1
	for base <= top {
		middle := base + (top-base)/2
		if target < nums[middle] {
			top = middle - 1
		} else if target > nums[middle] {
			base = middle + 1
		} else {
			return middle
		}
	}
	return base
}

// O(n) solution
// func searchInsert(nums []int, target int) int {
// 	if nums[0] >= target {
// 		return 0
// 	}
// 	if nums[len(nums)-1] < target {
// 		return len(nums)
// 	}
// 	for i, num := range nums {
// 		if num == target {
// 			return i
// 		} else if num > target {
// 			return i
// 		}
// 	}
// 	return len(nums)
// }

func main() {
	nums := []int{1, 3}
	target := 3
	fmt.Println(searchInsert(nums, target))
}
