/**
 *和为K的子数组
 * 
 */

func subarraySum(nums []int, k int) int {
    // 解法一：暴力枚举
    // ans := 0
    // for i:=0;i<len(nums);i++{
    //     cen := 0
    //     for j:=i;j>=0;j--{
    //         cen += nums[j]
    //         if cen == k {
    //             ans ++
    //         }
    //     }
    // }
    // return ans

    // 解法二：前缀数组+map

    m := map[int]int{}
    pre,count:=0,0
    m[0] = 1
    for i:=0;i<len(nums);i++{
        pre += nums[i]
        if _,ok := m[pre-k];ok {
            count += m[pre-k]
        }

        m[pre] ++
    }

    return count
}