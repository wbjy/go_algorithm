/** 
 *单词拆分 II
 *
 * 解题思路是动态规划+回朔法
 *
 * 用动态规划来判断字符串能不能拆分，用回溯法来遍历所有的结果
 * 
 */
func wordBreak(s string, wordDict []string) []string {
    m :=make(map[string]bool, len(wordDict))
    for _,v := range wordDict {
        m[v] = true
    }
    
    dp := make([]bool, len(s)+1)
    dp[0] = true
    for i:=1;i<=len(s);i++{
        for j:=0;j<i;j++{
            if dp[j]&&m[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    str := []string{}
    if !dp[len(s)]{
        return str 
    }
    
    var dfs = func(string, []string){}
    dfs = func(ns string, r []string) {
        if len(ns) == 0{
            str = append(str, strings.Join(r," "))
            return
        }
        for i:=1;i<=len(ns);i++{
            if m[ns[:i]] {
                dfs(ns[i:], append(r, ns[:i]))
            }
        }
    }
    
    dfs(s, []string{})
    
    return str
    
}