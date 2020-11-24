/** 
 *最长上升子序列
 *
 * 
 */
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = 1
	
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}

	max := 0
	for i := 0; i < n; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}