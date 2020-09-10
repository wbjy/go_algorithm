/**
 *打家劫舍
 *
 * 
 */
func rob(nums []int) int {
    if len(nums) == 0{
     return 0   
    }
    if len(nums) == 1{
        return nums[0]
    }
    
    first := nums[0]
    sec := max(nums[0], nums[1])
    for i:=2;i<len(nums);i++{
        first,sec = sec,max(first+nums[i], sec)
    }
    
    return sec
}

func max(x, y int) int{
    if x > y {
        return x
    }
    return y
}