// 1. Two Sum
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

package main

import "fmt"

func main() {

	target := 6
	nums := []int{3, 2, 3}
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Println(nums[i], nums[j])
				return []int{i, j}
			}
		}
	}

	return []int{}

}
