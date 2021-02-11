/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i] + 1
		if d%10 == 0 {
			digits[i] = 0
			continue
		}
		digits[i] = d
		return digits
	}
	newDigits := []int{1}
	for i := 0; i < len(digits); i++ {
		newDigits = append(newDigits, 0)
	}
	return newDigits
}

// @lc code=end
