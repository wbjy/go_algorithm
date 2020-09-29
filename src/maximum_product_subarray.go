/** 
 *乘积最大子数组
 * 
 */
func maxProduct(nums []int) int {
    maxInt := nums[0]
    minInt := nums[0]
    ans := nums[0]
    for i:=1;i<len(nums);i++{
        //解法一
        // mx,mn := maxInt,minInt
        // maxInt = max(nums[i]*mx, max(nums[i],nums[i]*mn))
        // minInt = min(nums[i]*mn, min(nums[i], nums[i]*mx))
        // ans = max(ans, maxInt)
        //解法二
        mx,mn := maxInt,minInt
        if nums[i] > 0{
            maxInt = max(nums[i], mx*nums[i])
            minInt = min(nums[i], mn*nums[i])
        }else{
            maxInt = max(nums[i], mn*nums[i])
            minInt = min(nums[i], mx*nums[i])
        }
        ans = max(ans, maxInt)
        
    }
    
    return ans
}

func min(a, b int) int {
    if a>b{
        return b
    }
    return a
}

func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}