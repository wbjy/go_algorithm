/** 
 * 缺失数字
 */

func missingNumber(nums []int) int {
    total := 0
    
    for i:=0;i<len(nums);i++{
        total += nums[i]
    }
    maxNum := len(nums)
    return (maxNum+1)*maxNum/2-total
}