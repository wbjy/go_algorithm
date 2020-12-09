/**
 *回文子串
 * 
 */
// 解法一：暴力解法
// func countSubstrings(s string) int {
//     if len(s) <= 1 {
//         return len(s)
//     }
//     ans := len(s)
//     for i:=1;i<len(s);i++{
//         for j:=i-1;j>=0;j--{
//             if test(s[j:i+1]) {
//                 ans ++
//             }
//         }
//     }

//     return ans
// }

// func test(str string) bool{
//     start, end := 0, len(str)-1
//     for start < end {
//         if str[start] != str[end] {
//             return false
//         }else {
//             start ++
//             end --
//         }
//     }

//     return true
// }

// 解法二：动态规划
func countSubstrings(s string) int {
    if len(s) <= 1 {
        return len(s)
    }
    dp := make([]bool, len(s))
    count := 0
    for i:=0;i<len(s);i++{
        for j:=0;j<=i;j++{
            if i == j {
                count ++
                dp[j] = true
            }else if i-j == 1 && s[i] == s[j] {
                dp[j] = true
                count ++
            }else if i-j>1 && s[i] == s[j] && dp[j+1] {
                dp[j] = true
                count ++
            }else{
                dp[j] = false
            }
        }
    }
    return count

}