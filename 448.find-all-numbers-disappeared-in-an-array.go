/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 *
 * https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (59.35%)
 * Likes:    6984
 * Dislikes: 397
 * Total Accepted:    608.8K
 * Total Submissions: 1M
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * Given an array nums of n integers where nums[i] is in the range [1, n],
 * return an array of all the integers in the range [1, n] that do not appear
 * in nums.
 *
 *
 * Example 1:
 * Input: nums = [4,3,2,7,8,2,3,1]
 * Output: [5,6]
 * Example 2:
 * Input: nums = [1,1]
 * Output: [2]
 *
 *
 * Constraints:
 *
 *
 * n == nums.length
 * 1 <= n <= 10^5
 * 1 <= nums[i] <= n
 *
 *
 *
 * Follow up: Could you do it without extra space and in O(n) runtime? You may
 * assume the returned list does not count as extra space.
 *
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	buket := make([]int, len(nums))
	for _, num := range nums {
		buket[num-1] = num
	}
	j := 0
	for i := 0; i < len(buket); i++ {
		if buket[i] == 0 {
			buket[j] = i + 1
			j++
		}
	}
	return buket[:j]
}

// @lc code=end

