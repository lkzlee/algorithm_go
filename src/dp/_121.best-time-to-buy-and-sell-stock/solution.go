package _121_best_time_to_buy_and_sell_stock

/***
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

注意：你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//解法2，更优的dp问题，只需要常数空间，在解法1基础上优化
func MaxProfit2(prices []int) int {
	//双层循环，最笨的办法找出最大收益，暴力破解
	//第二种解法，此动态规划不是直接找解，而是记录每次最小的买入价
	//dp[i] 表示0...i的最低股票值
	if len(prices) <= 0 {
		return 0
	}
	minValue := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-minValue > maxProfit {
			maxProfit = prices[i] - minValue
		}
		if prices[i] < minValue {
			minValue = prices[i]
		}
	}
	return maxProfit

}

//解法1，普通dp问题，需要o（n）的空间复杂度
func MaxProfit(prices []int) int {
	//双层循环，最笨的办法找出最大收益，暴力破解
	//第二种解法，此动态规划不是直接找解，而是记录每次最小的买入价
	//dp[i] 表示0...i的最低股票值
	if len(prices) <= 0 {
		return 0
	}
	dp := make([]int, len(prices))
	dp[0] = prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < dp[i-1] {
			dp[i] = prices[i]
		} else {
			dp[i] = dp[i-1]
		}
		if prices[i]-dp[i-1] > maxProfit {
			maxProfit = prices[i] - dp[i-1]
		}
	}
	return maxProfit

}
