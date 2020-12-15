/**
 *跳跃游戏
 * 
 */
func canJump(nums []int) bool {
    i := 0
    next := 0
    for i < len(nums) {
        if next >= len(nums) -1 {
            return true
        }
        if nums[i] != 0{
            next = max(next, nums[i] + i)
            i++
        }else {
            if i == next {
                return false
            }
            i ++
        }
    }

    return true
}

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}