/** 
 *最长有效括号
 *
 * 解题思路：看到这题，第一眼想到用栈的先进后出规则处理这没问题，难点在于arr = append(arr, -1)提前存入一个-1这个操作
 *           
 * 
 */
func longestValidParentheses(s string) int {
    arr := make([]int,0)
    arr = append(arr, -1)
    maxNum := 0
    for i:=0;i<len(s);i++{
       if s[i] == '(' {
           arr = append(arr, i)
       }else{
           arr = arr[:len(arr)-1]
           if len(arr) == 0{
               arr = append(arr, i)
           }else{
               maxNum = max(maxNum, i - arr[len(arr)-1])
           }
       }
    }
    return maxNum
}

func max(a,b int) int{
    if a > b {
        return a
    }
    return b
}