/*
 * @lc app=leetcode id=1346 lang=golang
 *
 * [1346] Check If N and Its Double Exist
 *
 * https://leetcode.com/problems/check-if-n-and-its-double-exist/description/
 *
 * algorithms
 * Easy (35.85%)
 * Likes:    1093
 * Dislikes: 132
 * Total Accepted:    218.2K
 * Total Submissions: 608.7K
 * Testcase Example:  '[10,2,5,3]'
 *
 * Given an array arr of integers, check if there exists two integers N and M
 * such that N is the double of M ( i.e. N = 2 * M).
 *
 * More formally check if there exists two indices i and j such that :
 *
 *
 * i != j
 * 0 <= i, j < arr.length
 * arr[i] == 2 * arr[j]
 *
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [10,2,5,3]
 * Output: true
 * Explanation: N = 10 is the double of M = 5,that is, 10 = 2 * 5.
 *
 *
 * Example 2:
 *
 *
 * Input: arr = [7,1,14,11]
 * Output: true
 * Explanation: N = 14 is the double of M = 7,that is, 14 = 2 * 7.
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [3,1,7,11]
 * Output: false
 * Explanation: In this case does not exist N and M, such that N = 2 * M.
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= arr.length <= 500
 * -10^3 <= arr[i] <= 10^3
 *
 *
 */

// @lc code=start
func checkIfExist(arr []int) bool {
	exist := make(map[int]bool)
	for _, v := range arr {
		if exist[2*v] || (v%2 == 0 && exist[v/2]) {
			return true
		}
		exist[v] = true
	}
	return false
}

// @lc code=end

