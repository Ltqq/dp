package kadane

//-   [LeetCode 918: Maximum Sum Circular Subarray](https://leetcode.com/problems/maximum-sum-circular-subarray/description/)
//给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
//
// 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个
//元素是 nums[(i - 1 + n) % n] 。
//
// 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不
//存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,-2,3,-2]
//输出：3
//解释：从子数组 [3] 得到最大和 3
//
//
// 示例 2：
//
//
//输入：nums = [5,-3,5]
//输出：10
//解释：从子数组 [5,5] 得到最大和 5 + 5 = 10
//
//
// 示例 3：
//
//
//输入：nums = [3,-2,2,-3]
//输出：3
//解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 3 * 10⁴
// -3 * 10⁴ <= nums[i] <= 3 * 10⁴
//
//
// Related Topics 队列 数组 分治 动态规划 单调队列
// 👍 718 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
// 1.子问题: dp[i] = nums[(i - 1 + n) % n] +  nums[i]  nums[(i - 1 + n) % n] > 0
//
//	dp[i] = nums[i]
func maxSubarraySumCircular(nums []int) int {
	return 0
}
