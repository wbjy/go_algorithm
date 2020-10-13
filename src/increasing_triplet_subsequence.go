/** 
 *递增的三元子序列
 *
 * 解法一：顺序遍历出最小的m1和次小的m2,然后大于m1和m2的数就满足条件
 * 解法二：动态规划，思路是找一个数组dp[i]分别顺序存储nums[i]大于之前nums[0---i]的个数,当个数大于等于3就满足条件
 * 
 */
func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {
        return false
    }
    // 解法一
    // m1,m2 := math.MaxInt32,math.MaxInt32
    // for _,v :=range nums{
    //     if m1 >= v {
    //         m1 = v
    //     }else if m2 >= v {
    //         m2 = v
    //     }else {
    //         return true
    //     }
    // }
    // return false
    // 
    //解法二 动态规划
    dp := make([]int, len(nums))
    for e := range nums{
        dp[e] = 1
    }
    
    for i:=1;i<len(nums);i++{
        for j:=0;j<i;j++{
            if nums[i] > nums[j]{
                dp[i] = max(dp[i], dp[j]+1)
            }
            if dp[i] >=3{
                return true
            }
        }
    }
    
    return false
    
}

func max(a, b int) int{
    if a > b{
        return a
    }
    
    return b
}