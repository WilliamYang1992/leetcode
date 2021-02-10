/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	l := len(nums)
	left := 0
	right := l - 1
	ret := l
	for left <= right {
		mid := (right-left)>>1 + left
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			ret = mid
			right--
		} else {
			left++
		}
	}
	return ret
}

// @lc code=end
