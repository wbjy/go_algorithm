/**
 *分割等和子集
 *解题思路：动态规划---> 典型的0-1背包问题
 * 
 */
func canPartition(nums []int) bool {
    n := len(nums)
    if n < 2 {
        return false
    }
    total := 0
    max := 0
    for _,v := range nums {
        total += v
        if max < v {
            max = v
        }
    }
    if total % 2 != 0 {
        return false
    }
    target := total/2
    if max > target {
        return false
    }
    
    
    // dp := make([][]bool, n)
    // for i := range dp {
    //     dp[i] = make([]bool, target+1)
    // }
    // for i:=0;i<n;i++{
    //     dp[i][0] = true
    // }
    // dp[0][nums[0]] = true
    // for i:=1;i<n;i++{
    //     v := nums[i]
    //     for j:=1;j<= target;j++{
    //         if j >= v{
    //             dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
    //         }else{
    //             dp[i][j] = dp[i-1][j]
    //         }
    //     }
    // }
    // return dp[n-1][target]
    dp := make([]bool, target+1)
    dp[0] = true
    for i:=0;i<n;i++{
        v := nums[i]
        for j:= target;j>=v;j--{
            dp[j] = dp[j]||dp[j-v]
        }
    }
    return dp[target]
}