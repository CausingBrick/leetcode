/*
 * @lc app=leetcode id=941 lang=golang
 *
 * [941] Valid Mountain Array
 *
 * https://leetcode.com/problems/valid-mountain-array/description/
 *
 * algorithms
 * Easy (33.64%)
 * Likes:    2232
 * Dislikes: 147
 * Total Accepted:    310.8K
 * Total Submissions: 924K
 * Testcase Example:  '[2,1]'
 *
 * Given an array of integers arr, return true if and only if it is a valid
 * mountain array.
 *
 * Recall that arr is a mountain array if and only if:
 *
 *
 * arr.length >= 3
 * There exists some i with 0 < i < arr.length - 1 such that:
 *
 * arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
 * arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
 *
 *
 *
 *
 *
 * Example 1:
 * Input: arr = [2,1]
 * Output: false
 * Example 2:
 * Input: arr = [3,5,5]
 * Output: false
 * Example 3:
 * Input: arr = [0,3,2,1]
 * Output: true
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 10^4
 * 0 <= arr[i] <= 10^4
 *
 *
 */

// @lc code=start
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	lo, hi := 0, len(arr)-1
	for lo+1 < len(arr) {
		if arr[lo+1] <= arr[lo] {
			break
		}
		lo++
	}
	for hi-1 >= 0 {
		if arr[hi-1] <= arr[hi] {
			break
		}
		hi--
	}
	return lo == hi && lo != 0 && hi != len(arr)-1
}

// @lc code=end

