/**
 *每日温度
 *
 * 解题思路：像这种找后面第一个最大值的提可以考虑用单调栈
 * 
 */
func dailyTemperatures(T []int) []int {
    res := make([]int, len(T))
    stack := make([]int, 0)
    for i:=0;i<len(T);i++{
        t := T[i]
       for len(stack) > 0 && t > T[stack[len(stack)-1]] {
           index := stack[len(stack)-1]
           stack = stack[:len(stack)-1]
           res[index] = i-index
       }
       stack = append(stack, i)
    }
    return res
}