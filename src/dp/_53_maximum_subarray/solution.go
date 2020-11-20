package _53_maximum_subarray

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func MaxSubArray(nums []int) int {
	//这一题可以用 DP 求解也可以不用 DP。
	//题目要求输出数组中某个区间内数字之和最大的那个值。
	//dp[i] 表示 [0,i] 区间内各个子区间和的最大值，
	//状态转移方程是:
	//dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0)
	//dp[i] = nums[i] (dp[i-1] ≤ 0)。
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
	}
	max := dp[0]
	for i := 1; i < len(dp); i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
