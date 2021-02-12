/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	ret := make([][]int, numRows)
	ret[0] = []int{1}
	for i := 1; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		ret[i][0] = 1
		ret[i][i] = 1
		for j := 1; j < i; j++ {
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
	}
	return ret
}

// @lc code=end
