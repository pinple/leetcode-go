/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	n := 0
	for x != 0 {
		if n < math.MinInt32/10 || n > math.MaxInt32/10 {
			return 0
		}
		n = n*10 + x%10
		x = x / 10
	}
	return n
}

// @lc code=end

