/** 
 *无重复字符的最长子串
 *解题思路：滑动窗口
 * 
 */

func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    dp := make([]int, len(s))
    dp[0] = 1
    for i:=1;i<len(s);i++{
        dp[i] = 1
        for j:=i-dp[i-1];j<i;j++{
            if s[j] == s[i] {
                dp[i] = 1
            }else{
                dp[i] ++
            }
        }
    }
    max := 1
    for _, v := range dp{
        if max < v {
            max = v
        }
    }

    return max

}