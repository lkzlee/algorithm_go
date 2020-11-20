package _0063_unique_paths_ii

/***
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？



网格中的障碍物和空位置分别用 1 和 0 来表示。



示例 1：


输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
示例 2：


输入：obstacleGrid = [[0,1],[0,0]]
输出：1


提示：
m == obstacleGrid.length
n == obstacleGrid[i].length
1 <= m, n <= 100
obstacleGrid[i][j] 为 0 或 1
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func UniquePathsWithObstacles(og [][]int) int {
	//这一题是加强版的dp问题，相比较地62题多了一个限定条件是否有障碍物
	//状态转移方程为: dp[i][j]=dp[i-1][j]+dp[i][j-1]
	//og[i-1][j]==1 时dp[i][j]=0
	dp := make([][]int, len(og))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(og[i]))
	}
	if og[0][0] == 1 {
		return 0
	}
	dp[0][0] = 1
	//初始化行边
	for i := 1; i < len(og); i++ {
		if og[i][0] == 1 {
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	//初始化列边
	for j := 1; j < len(og[0]); j++ {
		if og[0][j] == 1 {
			dp[0][j] = 0
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < len(og); i++ {
		for j := 1; j < len(og[i]); j++ {
			if og[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	//fmt.Println(dp)
	return dp[len(dp)-1][len(dp[0])-1]
}
