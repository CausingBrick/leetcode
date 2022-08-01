/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 *
 * https://leetcode.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (32.14%)
 * Likes:    1838
 * Dislikes: 2408
 * Total Accepted:    349.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,2,1]'
 *
 * Given an integer array nums, return the third distinct maximum number in
 * this array. If the third maximum does not exist, return the maximum
 * number.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [3,2,1]
 * Output: 1
 * Explanation:
 * The first distinct maximum is 3.
 * The second distinct maximum is 2.
 * The third distinct maximum is 1.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2]
 * Output: 2
 * Explanation:
 * The first distinct maximum is 2.
 * The second distinct maximum is 1.
 * The third distinct maximum does not exist, so the maximum (2) is returned
 * instead.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [2,2,3,1]
 * Output: 1
 * Explanation:
 * The first distinct maximum is 3.
 * The second distinct maximum is 2 (both 2's are counted together since they
 * have the same value).
 * The third distinct maximum is 1.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 *
 * Follow up: Can you find an O(n) solution?
 */

// @lc code=start
func thirdMax(nums []int) int {
	max := [3]int{-1 << 32, -1 << 32, -1 << 32}
	for _, v := range nums {
		if v == max[0] || v == max[1] || v == max[2] {
			continue
		}
		if v > max[0] {
			max[2] = max[1]
			max[1] = max[0]
			max[0] = v
		} else if v > max[1] {
			max[2] = max[1]
			max[1] = v
		} else if v > max[2] {
			max[2] = v
		}
	}
	if max[2] == -1<<32 {
		return max[0]
	}
	return max[2]
}

// @lc code=end

