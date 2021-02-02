package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	subs := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]
		if _, ok := subs[sub]; ok {
			return []int{subs[sub], i}
		}
		subs[nums[i]] = i
	}
	return nil
}

// @lc code=end
