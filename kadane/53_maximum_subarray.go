package kadane

// Kadane's Algorithm is primarily used for solving the Maximum Subarray Problem
// and its variations where the problem asks
// to optimize a contiguous subarray within a one-dimensional numeric array
//
// #### LeetCode Problems:
//
// -   [LeetCode 53: Maximum Subarray](https://leetcode.com/problems/maximum-subarray/description/
func maxSubArray(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > result {
			result = nums[i]
		}
	}
	return result
}

// 1.找到子问题，前i个结尾的最大值 = 前i-1个结尾的最大值 + nums[i]   前i-1个结尾的最大值 > 0
// or 前i个结尾的最大值 = nums[i] 0
// 2.初始值 i=0 = nums[0]
func maxSubArray2(nums []int) int {
	var result = nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

//-   [LeetCode 152: Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/description/)
