/*
 * @lc app=leetcode id=1790 lang=golang
 *
 * [1790] Check if One String Swap Can Make Strings Equal
 *
 * https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/description/
 *
 * algorithms
 * Easy (45.63%)
 * Likes:    585
 * Dislikes: 26
 * Total Accepted:    53.1K
 * Total Submissions: 116.5K
 * Testcase Example:  '"bank"\n"kanb"'
 *
 * You are given two strings s1 and s2 of equal length. A string swap is an
 * operation where you choose two indices in a string (not necessarily
 * different) and swap the characters at these indices.
 *
 * Return true if it is possible to make both strings equal by performing at
 * most one string swap on exactly one of the strings. Otherwise, return
 * false.
 *
 *
 * Example 1:
 *
 *
 * Input: s1 = "bank", s2 = "kanb"
 * Output: true
 * Explanation: For example, swap the first character with the last character
 * of s2 to make "bank".
 *
 *
 * Example 2:
 *
 *
 * Input: s1 = "attack", s2 = "defend"
 * Output: false
 * Explanation: It is impossible to make them equal with one string swap.
 *
 *
 * Example 3:
 *
 *
 * Input: s1 = "kelb", s2 = "kelb"
 * Output: true
 * Explanation: The two strings are already equal, so no string swap operation
 * is required.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s1.length, s2.length <= 100
 * s1.length == s2.length
 * s1 and s2 consist of only lowercase English letters.
 *
 *
 */

// @lc code=start
func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	c1, c2 := [26]int{}, [26]int{}
	diffCount := 0
	for i := 0; i < len(s1); i++ {
		c1[s1[i]-97]++
		c2[s2[i]-97]++
		if s1[i] != s2[i] {
			diffCount++
		}
	}
	if c1 == c2 && diffCount == 2 {
		return true
	}
	return false
}

// @lc code=end

