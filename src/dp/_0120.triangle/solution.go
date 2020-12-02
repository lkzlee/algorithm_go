package _0120_triangle

import "math"

/***
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。



例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。



说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//解法3 最好的解法，倒序dp处理，无需空间
func MinimumTotal3(triangle [][]int) int {
	if len(triangle) <= 0 {
		return 0
	}
	for i := len(triangle) - 1; i > 0; i-- {
		for j := len(triangle[i]) - 2; j >= 0; j-- {
			triangle[i-1][j] += Min(triangle[i][j], triangle[i][j+1])
		}
	}
	return triangle[0][0]
}

//解法2 更优的动态规划方法，需要空间一维数组dp
func MinimumTotal2(triangle [][]int) int {
	if len(triangle) <= 0 {
		return 0
	}
	//动态规划转移方程
	//dp[j]=min(dp[j],dp[j-1])+triangle[i][j]
	dp := make([]int, len(triangle[len(triangle)-1]))
	for i := 0; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if i == 0 {
				dp[j] = triangle[i][j]
				continue
			}
			//处理最左边
			if j == 0 {
				dp[j] = dp[j] + triangle[i][j]
				continue
			}
			//处理最右边
			if j == len(triangle[i-1]) {
				dp[j] = dp[j-1] + triangle[i][j]
				continue
			}
			dp[j] = Min(dp[j], dp[j-1]) + triangle[i][j]
		}
	}
	minNum := math.MaxInt64
	for j := 0; j < len(dp); j++ {
		minNum = Min(minNum, dp[j])
	}
	return minNum

}

//解法1 普通的动态规划方法，需要空间二维数组dp
func MinimumTotal(triangle [][]int) int {
	if len(triangle) <= 0 {
		return 0
	}
	//动态规划转移方程
	//dp[i][j]=min(dp[i-1][j],dp[i-1][j-1])+triangle[i][j]
	dp := make([][]int, len(triangle))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if i == 0 {
				dp[i][j] = triangle[i][j]
				continue
			}
			//处理当前最左
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
				continue
			}
			//处理最右
			if j == len(dp[i-1]) {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
				continue
			}
			dp[i][j] = Min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
	}
	min := dp[len(triangle)-1][0]
	for j := 0; j < len(triangle[len(triangle)-1]); j++ {
		min = Min(min, dp[len(triangle)-1][j])
	}
	return min
}

func Min(m1 int, m2 int) int {
	if m1 < m2 {
		return m1
	}
	return m2
}
