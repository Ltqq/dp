package kadane

// Kadane's Algorithm is primarily used for solving the Maximum Subarray Problem
// and its variations where the problem asks
// to optimize a contiguous subarray within a one-dimensional numeric array
//
// #### LeetCode Problems:
//
// -   [LeetCode 53: Maximum Subarray](https://leetcode.com/problems/maximum-subarray/description/
// TODO understand why
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//-   [LeetCode 918: Maximum Sum Circular Subarray](https://leetcode.com/problems/maximum-sum-circular-subarray/description/)
//-   [LeetCode 152: Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/description/)
