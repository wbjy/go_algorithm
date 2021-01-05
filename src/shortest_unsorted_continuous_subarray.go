/**
 *最短无序连续子数组
 * 
 */
func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return 0
    }
    l, r := 0, n-1
    for l < n-1 &&  nums[l] <= nums[l+1] {
        l ++
    }
    for r > 0 && nums[r] >= nums[r-1] {
        r --
    }

    if l >= r {
        return 0
    }
    min, max := getBroder(nums, l, r)
    for l > 0 && nums[l-1] > min {
        l --
    }

    for r < n-1 && nums[r+1] < max {
        r ++
    }
    return r-l+1

}

func getBroder(nums []int, l, r int) (int, int) {
    min, max := nums[l], nums[r] 
    for i := l;i <= r; i++{
        if nums[i] > max {
            max = nums[i]
        }
        if nums[i] < min {
            min = nums[i]
        }
    }

    return min, max
}