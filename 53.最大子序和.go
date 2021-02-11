/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{
			nums[l], nums[l], nums[l], nums[l],
		}
	}
	m := (r-l)>>1 + l
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func pushUp(l, r Status) Status {
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	iSum := l.iSum + r.iSum
	return Status{lSum, rSum, mSum, iSum}
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

type Status struct {
	lSum, rSum, mSum, iSum int
}

// @lc code=end
