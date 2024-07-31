package kadane

//Kadane's Algorithm is primarily used for solving the Maximum Subarray Problem
//and its variations where the problem asks
//to optimize a contiguous subarray within a one-dimensional numeric array
//
//#### LeetCode Problems:
//
//-   [LeetCode 53: Maximum Subarray](https://leetcode.com/problems/maximum-subarray/description/
//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。
//
//
//
// 示例 1：
//
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组[4,-1,2,1] 的和最大，为6 。
//
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：1
//
//
// 示例 3：
//
//
//输入：nums = [5,4,-1,7,8]
//输出：23
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
//
// Related Topics 数组 分治 动态规划
// 👍 6725 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

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

//leetcode submit region end(Prohibit modification and deletion)

//-   [LeetCode 918: Maximum Sum Circular Subarray](https://leetcode.com/problems/maximum-sum-circular-subarray/description/)
//-   [LeetCode 152: Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/description/)
