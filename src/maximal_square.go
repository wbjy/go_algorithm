/**
 *最大正方形
 *
 * 解题思路：动态规划
 */

func maximalSquare(matrix [][]byte) int {
    dp := make([][]int, len(matrix))
    maxSize := 0
    for i:=0; i<len(matrix); i++{
        dp[i] = make([]int, len(matrix[i]))
        for j := 0; j < len(matrix[i]); j++ {
            dp[i][j] = int(matrix[i][j] - '0')
            if dp[i][j] == 1{
                maxSize = 1
            }
        }
    }
    
    for i := 1; i <len(matrix); i++ {
        for j := 1; j < len(matrix[i]); j++ {
            if dp[i][j] == 1 {
                dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
                
                if dp[i][j] > maxSize {
                    maxSize = dp[i][j]
                }
            }
        }
    }
    
    return maxSize*maxSize
}

func min(x, y int) int {
    if x < y {
        return x
    }
    
    return y
}

