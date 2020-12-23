/**
 *编辑距离
 *
 * 解题思路：万能的动态规划
 * 
 */
func minDistance(word1 string, word2 string) int {
    n,m := len(word1),len(word2)
    var dp = make([][]int, n+1)
    for i:=0;i<= n;i++{
        dp[i] = make([]int, m+1)
    }
    if n*m == 0 {
        return n+m
    }
	for i := 0; i <= n; i++ {
		dp[i][0] = i 
	}
	// dp[0][i]表示将空串编辑为str2[:i]的代价，即插入字符的代价
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}
	// 下面是动态规划的主方法
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]              // 如果 word1[i-1] 与 word2[i-1]相等
			} else {
				dp[i][j] = dp[i-1][j-1] + 1          // 替换代价
			}
			dp[i][j] = min(dp[i][j], dp[i][j-1] + 1) // 插入代价
			dp[i][j] = min(dp[i][j], dp[i-1][j] + 1) // 删除代价
		}
	}
	return dp[n][m]
}

func min(a,b int)int{
    if a < b {
        return a
    }
    return b
}