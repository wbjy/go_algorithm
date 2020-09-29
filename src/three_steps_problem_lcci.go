/**
 *三步问题
 *
 * 动态规划
 * 
 */

func waysToStep(n int) int {
//  解法一 效率有点低
//     arr := make([]int, n+1)
//     arr[0] = 1
//     for i:=1;i<=n;i++{
//         if i==1{
//             arr[i] = 1
//         }else if i == 2{
//             arr[i] = (arr[i-1]+arr[i-2]) % 1000000007
//         }else {
//             arr[i] = (arr[i-1]+arr[i-2]+arr[i-3]) % 1000000007
//         }
//     }
    
//     return arr[n]
//     解法二 由于只需要保存前三个数就可以了，所以不需要数组来存储数据，只需要保留i-3,i-2,i-1
    
    if n==0{
        return 0
    }
    if n==1{
        return 1
    }
    if n==2{
        return 2
    }
    if n==3{
        return 4
    }
    
    dp1,dp2,dp3:=1,2,4
    
    for i:=4;i<=n;i++{
        dp1,dp2,dp3 = dp2,dp3,(dp1+dp2+dp3)%1000000007
    }
    
    return dp3
    
}