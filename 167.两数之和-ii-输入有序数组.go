/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		low := i + 1
		high := len(numbers) - 1
		for low <= high {
			mid := (high-low)>>1 + low
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] < target-numbers[i] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return []int{-1, -1}
}

// @lc code=end
