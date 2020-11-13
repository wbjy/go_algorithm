/**
 *最长连续序列
 * 
 */
func longestConsecutive(nums []int) int {
    m := make(map[int]bool)
    for _,v := range nums {
        m[v] = true
    }

    longest := 0
    for num := range m {
        if !m[num-1]{ // 保证这个数是最开头的数字
            currtNum := num
            l := 1
            for m[currtNum+1] {
                l++
                currtNum++
            }
            if longest < l {
                longest = l
            }
        }
        
    }
    return longest
}