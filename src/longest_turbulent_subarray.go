/**
 *最长湍流子数组
 *
 * 
 */
func maxTurbulenceSize(arr []int) int {
    // 解法一： 动态规划
    // if len(arr) == 1 {
    //     return len(arr)
    // }
    // dp := make([]int, len(arr))
    // dp[0] = 1
    // flag := 0
    // for i := 1; i < len(arr); i ++ {
    //     if arr[i] > arr[i - 1]{
    //         if flag == 1 {
    //             dp[i] = 2
    //         }else if flag == -1 {
    //             dp[i] = dp[i-1] + 1
    //         }else{
    //             dp[i] = dp[i-1] +1
    //         }
    //         flag = 1
    //     }else if arr[i] < arr[i - 1]{
    //         if flag == 1{
    //             dp[i] = dp[i-1] + 1
    //         }else if flag == -1 {
    //             dp[i] = 2
    //         }else {
    //             dp[i] = dp[i-1] + 1
    //         }
    //         flag = -1
    //     }else{
    //         flag = 0
    //         dp[i] = 1
    //     }
    // }
    // max := dp[0]
    // for i := 1; i<len(dp); i ++ {
    //     if max < dp[i] {
    //         max = dp[i]
    //     }
    // }
    // return max

    // 解法二 ： 滑动窗口
    if len(arr) == 1 {
        return len(arr)
    }
    max := 1
    l := len(arr)
    i := 1
    for {
        for i < l && arr[i] == arr[i-1] {
            i++
        }
        if i == l {
            return max
        }
        ans,flag := 2, arr[i] > arr[i-1]
        i ++
        for i < l && arr[i] != arr[i-1] && (arr[i] > arr[i-1]) != flag {
            i ++
            ans ++
            flag = !flag
        }
        if ans > max {
            max = ans
        }
    }
    return max
}
