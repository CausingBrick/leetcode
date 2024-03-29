/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 *
 * https://leetcode.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (53.80%)
 * Likes:    6228
 * Dislikes: 792
 * Total Accepted:    894.7K
 * Total Submissions: 1.7M
 * Testcase Example:  '19'
 *
 * Write an algorithm to determine if a number n is happy.
 *
 * A happy number is a number defined by the following process:
 *
 *
 * Starting with any positive integer, replace the number by the sum of the
 * squares of its digits.
 * Repeat the process until the number equals 1 (where it will stay), or it
 * loops endlessly in a cycle which does not include 1.
 * Those numbers for which this process ends in 1 are happy.
 *
 *
 * Return true if n is a happy number, and false if not.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 19
 * Output: true
 * Explanation:
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 *
 *
 * Example 2:
 *
 *
 * Input: n = 2
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 2^31 - 1
 *
 *
 */

// @lc code=start
func isHappy(n int) bool {
	slow, fast := n, digitSquareSum(n)
	for slow != fast {
		slow = digitSquareSum(slow)
		fast = digitSquareSum(fast)
		fast = digitSquareSum(fast)
	}
	return slow == 1
}

func digitSquareSum(n int) int {
	sum := 0
	for n > 0 {
		mod := n % 10
		sum += mod * mod
		n /= 10
	}
	return sum
}

// @lc code=end

