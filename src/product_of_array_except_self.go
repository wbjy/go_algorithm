/** 
 *除自身以外数组的乘积
 *
 * 解题思路：用数组来保存res[i]左侧的乘积，然后在用total来保存res[i]右侧的乘积
 *
 *  时间复杂度 O(n)
 * 
 */

func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    res[0]=1
    for i:=1;i<len(nums);i++{
        res[i] = nums[i-1]*res[i-1]
    }
    
    total :=1
    for j:=len(nums)-1;j>=0;j--{
        res[j] = res[j]*total
        
        total= total*nums[j]
    }
    
    return res
}