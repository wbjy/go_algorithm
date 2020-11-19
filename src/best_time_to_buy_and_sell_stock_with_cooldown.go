/** 
 *最佳买卖股票时机含冷冻期
 *
 * 解题思路：动态规划
 * 
 */

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    dp := make([][3]int, len(prices))
    // 有股票
    dp[0][0] = -prices[0]
    // 无持有股票且处于冷冻期
    dp[0][1] = 0
    // 无持有股票且不处于冷冻期
    dp[0][2] = 0
    
    for i := 1; i < len(prices); i ++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][2] - prices[i])
        dp[i][1] = dp[i-1][0] + prices[i]
        dp[i][2] = max(dp[i-1][1], dp[i-1][2])
    }
    
    return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}