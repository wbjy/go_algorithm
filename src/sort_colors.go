/**
 *
 *颜色分类
 * 
 * 解题思路：双指针
 */
func sortColors(nums []int)  {
    // 解法一
    // left := 0
    // right := len(nums) -1
    // i := 0
    // for i < len(nums){
    //     if nums[i] == 0 {
    //         if i > left && nums[left] != 0 {
    //             nums[i],nums[left] = nums[left], nums[i]
    //             left ++
    //         }else if nums[left] == 0 && i > left{
    //             left ++
    //         }else{
    //             i++
    //         }
    //     }else if nums[i] == 2 {
    //         if i < right && nums[right] != 2{
    //             nums[i],nums[right] = nums[right], nums[i]
    //             right --
    //         }else if nums[right] == 2 && i < right{
    //             right --
    //         }else {
    //             break
    //         }
    //     } else{
    //         i ++
    //     }
    // }

    // 解法二
    a1,a2 := 0,0
    for k,v := range nums{
        if v == 0 {
            nums[a1],nums[k] = nums[k],nums[a1]
            if a1 < a2 {
                nums[a2],nums[k] = nums[k],nums[a2]
            }

            a1 ++
            a2 ++
        }else if v == 1 {
            nums[a2],nums[k] = nums[k], nums[a2]
            a2 ++
        }
    }
}