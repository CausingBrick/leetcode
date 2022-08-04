/*
 * @lc app=leetcode id=1502 lang=golang
 *
 * [1502] Can Make Arithmetic Progression From Sequence
 *
 * https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/description/
 *
 * algorithms
 * Easy (68.94%)
 * Likes:    761
 * Dislikes: 48
 * Total Accepted:    95.4K
 * Total Submissions: 138.5K
 * Testcase Example:  '[3,5,1]'
 *
 * A sequence of numbers is called an arithmetic progression if the difference
 * between any two consecutive elements is the same.
 *
 * Given an array of numbers arr, return true if the array can be rearranged to
 * form an arithmetic progression. Otherwise, return false.
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [3,5,1]
 * Output: true
 * Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with
 * differences 2 and -2 respectively, between each consecutive elements.
 *
 *
 * Example 2:
 *
 *
 * Input: arr = [1,2,4]
 * Output: false
 * Explanation: There is no way to reorder the elements to obtain an arithmetic
 * progression.
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= arr.length <= 1000
 * -10^6 <= arr[i] <= 10^6
 *
 *
 */

// @lc code=start
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

// @lc code=end

