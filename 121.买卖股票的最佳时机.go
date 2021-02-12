/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	var ret int
	var minValue int = 1e9
	for _, p := range prices {
		if p <= minValue {
			minValue = p
		} else if p-minValue > ret {
			ret = p - minValue
		}
	}
	return ret
}

// @lc code=end
