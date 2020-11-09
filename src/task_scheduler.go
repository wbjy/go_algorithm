/**
 *任务调度器
 *
 * 解题思路：主要是推导出数据公式 res = (maxNum - 1) * (n + 1) + len(maxCharArr), maxCharArr表示最大任务数组长度
 *
 * 
 */
func leastInterval(tasks []byte, n int) int {
    m := make(map[byte]int)
    maxNum := 0
    for _,v := range tasks {
        m[v] ++
    }
    for _,v := range m {
        if maxNum < v {
            maxNum = v
        }
    }
    res := 0
    res = (maxNum - 1) * (n + 1)
    for _,v := range m {
        if v == maxNum {
            res ++
        }
    }
    
    
    if len(tasks) > res {
        return len(tasks)
    }
    return res
}