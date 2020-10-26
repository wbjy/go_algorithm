/**
 *滑动窗口最大值
 *
 * 解题思路：主要是窗口数组，窗口数组里存最大值的坐标，并保证最左边的最大
 */
func maxSlidingWindow(nums []int, k int) []int {
    if nums == nil {
        return []int{}
    }
    res := make([]int, 0)
    //存下标
    win := make([]int, 0)
    
    for key,value := range nums {
        if key >= k && win[0] <= key - k{
            win = win[1:]
        }
        
        for len(win) > 0 && nums[win[len(win)-1]] <= value {
            win = win[:len(win)-1]
        }
        
        win = append(win, key)
        if key >= k-1{
            res = append(res, nums[win[0]])
        }
    }
    
    return res
    
}