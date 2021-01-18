/**
 *
 *零钱兑换
 *
 * 解题思路：这是个完全背包问题，也就是个典型的动态规划问题
 * 
 */
// func coinChange(coins []int, amount int) int {
//     dp := make([]int, amount+1)
//     dp[0] = 0
//     for k := range dp {
//         if k != 0{
//             dp[k] = amount + 1
//         }
//     }

//     for i := 1;i<=len(coins);i++{
//         for j:=0;j<=amount;j++{
//             if j >= coins[i-1]{
//                 dp[j] = min(dp[j], dp[j-coins[i-1]]+1)
//             }
//         }
//     }
//     if dp[amount] > amount {
//         return -1
//     }
//     return dp[amount]
// }

// func min(a, b int) int{
//     if a>b{
//         return b
//     }
//     return a
// }

// 未优化的动态规划解法
// func coinChange(coins []int, amount int) int {
//     dp := make([][]int, len(coins)+1)
//     for k := range dp {
//         dp[k] = make([]int, amount+1)
//     }

//     for i:=0;i<len(dp[0]);i++{
//         dp[0][i] = amount+1
//     }
//     dp[0][0] = 0
//     for i:=1;i<=len(coins);i++{
//         for j:=0;j<=amount;j++{
//             if j >= coins[i-1]{
//                 dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
//             }else{
//                 dp[i][j] = dp[i-1][j]
//             }
//         }
//     }
//     if dp[len(coins)][amount] > amount {
//         return -1
//     }
//     return dp[len(coins)][amount]
//     return 0
// }


//优化后的动态规划
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for k := range dp {
        dp[k] = amount + 1
    }
    dp[0] = 0
    for i:=1;i<=len(coins);i++{
        for j:=0;j<=amount;j++{
            if j >= coins[i-1] {
                dp[j] = min(dp[j], dp[j-coins[i-1]]+1)
            }
        }
    }
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}
func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}