/** 
 *一和零
 *
 * 解题思路：动态规划，主要是状态转移方程
 * 
 */
func findMaxForm(strs []string, m int, n int) int {
    dp := make([][]int, m+1)
    for index, _ := range dp {
        dp[index] = make([]int, n+1)
    }
    for _, s := range strs {
        zero := 0
        one := 0
        for _, ch := range s {
            if ch - '0' == 0 {
                zero ++
            }else {
                one ++
            }
        }
        
        for i := m; i >= zero; i--{
            for j := n; j >= one; j--{
                dp[i][j] = max(dp[i][j], 1 + dp[i-zero][j-one])
            }
        }
        
    }
    
    return dp[m][n]
    
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}