//1. Given an array of integers, return indices of the two numbers such that they add up to a specific target.
package main

import "fmt"

func calc(nums []int, target int) (int, int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target && nums[i] != nums[j] {
				return i, j
			}
		}
	}
	return -1, -1
}

func main() {
	// Change INPUT
	nums := []int{2, 7, 11, 15}
	target := 12
	//--------------------------
	i, j := calc(nums, target)
	if i == -1 {
		fmt.Printf("Not Found")
	} else {
		fmt.Printf("return[%d,%d]", i, j)
	}
}
