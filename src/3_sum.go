/**
 *三数之和
 *
 * 解题思路：先排序，在用双指针枚举
 * 
 */
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    ans := make([][]int, 0)
    n := len(nums)
    for first:=0;first<n;first++{
        if first >0 && nums[first] == nums[first-1]{
            continue
        }
        tir := n-1
        target := -1*nums[first]
        for sec := first + 1;sec<n;sec++{
            if sec > first+1 && nums[sec] == nums[sec-1]{
                continue
            }
            for sec < tir && nums[sec] + nums[tir] > target{
                tir--
            }
            if sec == tir {
                break
            }
            if nums[sec] + nums[tir] == target{
                ans = append(ans, []int{nums[first], nums[sec], nums[tir]})
            }
        }

    }
    return ans
}