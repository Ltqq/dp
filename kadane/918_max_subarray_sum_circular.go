package kadane

import "math"

// -   [LeetCode 918: Maximum Sum Circular Subarray](https://leetcode.com/problems/maximum-sum-circular-subarray/description/)
func maxSubarraySumCircular(nums []int) int {
	maxS := math.MinInt // 最大子数组和，不能为空
	minS := 0           // 最小子数组和，可以为空
	var maxF, minF, sum int
	for _, x := range nums {
		// 以 nums[i-1] 结尾的子数组选或不选（取 max）+ x = 以 x 结尾的最大子数组和
		maxF = max(maxF, 0) + x
		maxS = max(maxS, maxF)
		// 以 nums[i-1] 结尾的子数组选或不选（取 min）+ x = 以 x 结尾的最小子数组和
		minF = min(minF, 0) + x
		minS = min(minS, minF)
		sum += x
	}
	if sum == minS {
		return maxS
	}
	return max(maxS, sum-minS)
}
