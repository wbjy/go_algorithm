/**
 *找到所有数组中消失的数字
 *
 * 
 */
func findDisappearedNumbers(nums []int) []int {
    i := 0
    for i < len(nums) {
        if nums[i]-1 != i {
            j := nums[i]
            if nums[i] == nums[j-1] {
                nums[i] = 0
            }else{
                nums[i],nums[j-1] = nums[j-1],nums[i]
            }
        }
        if nums[i]-1 == i || nums[i] == 0{
            i ++
        }
        // fmt.Println(nums, i)
    }
    
    res := []int{}
    for k,v := range nums {
        if v == 0 {
            res = append(res, k+1)
        }
    }
    return res
}