/**
 *目标和
 *
 * 解题思路：转化一下之后就成了一个背包问题了
 * 
 */
func findTargetSumWays(nums []int, S int) int {
/**
    SumAdd + SumCut = sum
    SumAdd- SumCut = S
    2 * SumAdd = sum + S
    SumAdd = (sum+S)/2

    到这里就转化成nums有多少种加成SumAdd的情况
*/
    sum := 0
    for _, v := range nums {
        sum += v
    }

    if sum < S {
        return 0
    }

    if (sum+S) & 1 != 0{
        return 0
    }
    sumAdd := (sum+S)/2

    dp := make([]int, sumAdd+1)
    dp[0] = 1
    for _, num := range nums {
        for i := sumAdd;i >= num;i -- {
            dp[i] += dp[i-num]
        }
    } 
    return dp[sumAdd]
}