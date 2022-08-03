/*
 * @lc app=leetcode id=976 lang=golang
 *
 * [976] Largest Perimeter Triangle
 *
 * https://leetcode.com/problems/largest-perimeter-triangle/description/
 *
 * algorithms
 * Easy (53.46%)
 * Likes:    1276
 * Dislikes: 182
 * Total Accepted:    99.2K
 * Total Submissions: 186.1K
 * Testcase Example:  '[2,1,2]'
 *
 * Given an integer array nums, return the largest perimeter of a triangle with
 * a non-zero area, formed from three of these lengths. If it is impossible to
 * form any triangle of a non-zero area, return 0.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,1,2]
 * Output: 5
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,1]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= nums.length <= 10^4
 * 1 <= nums[i] <= 10^6
 *
 *
 */

// @lc code=start
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i > 1; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

// @lc code=end

