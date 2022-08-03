/*
 * @lc app=leetcode id=1491 lang=golang
 *
 * [1491] Average Salary Excluding the Minimum and Maximum Salary
 *
 * https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/description/
 *
 * algorithms
 * Easy (63.90%)
 * Likes:    809
 * Dislikes: 101
 * Total Accepted:    130.7K
 * Total Submissions: 205K
 * Testcase Example:  '[4000,3000,1000,2000]'
 *
 * You are given an array of unique integers salary where salary[i] is the
 * salary of the i^th employee.
 *
 * Return the average salary of employees excluding the minimum and maximum
 * salary. Answers within 10^-5 of the actual answer will be accepted.
 *
 *
 * Example 1:
 *
 *
 * Input: salary = [4000,3000,1000,2000]
 * Output: 2500.00000
 * Explanation: Minimum salary and maximum salary are 1000 and 4000
 * respectively.
 * Average salary excluding minimum and maximum salary is (2000+3000) / 2 =
 * 2500
 *
 *
 * Example 2:
 *
 *
 * Input: salary = [1000,2000,3000]
 * Output: 2000.00000
 * Explanation: Minimum salary and maximum salary are 1000 and 3000
 * respectively.
 * Average salary excluding minimum and maximum salary is (2000) / 1 = 2000
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= salary.length <= 100
 * 1000 <= salary[i] <= 10^6
 * All the integers of salary are unique.
 *
 *
 */

// @lc code=start
func average(salary []int) float64 {
	min, max := 1<<63-1, -1<<63
	sum := 0
	for _, s := range salary {
		if s > max {
			max = s
		}
		if s < min {
			min = s
		}
		sum += s
	}
	return float64(sum-max-min) / float64(len(salary)-2)
}

// @lc code=end

