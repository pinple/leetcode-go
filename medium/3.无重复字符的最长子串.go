/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start := 0
	n := 0
	for i, c := range s {
		if v, ok := m[c]; ok {
			start = max(v, start)
		}
		m[c] = i + 1
		n = max(i-start+1, n)
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

