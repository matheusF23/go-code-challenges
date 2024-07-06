package main

import "fmt"

func removeElement(nums []int, val int) int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		}
	}
	fmt.Println(nums)
	return l
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
