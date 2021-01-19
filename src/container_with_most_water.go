/**
 *盛最多水的容器
 *
 * 解题思路：双指针手法，双指针因为长度越往中间肯定是越小，所以left和right谁小谁往中间移动
 * 
 */

func maxArea(height []int) int {
    maxNum := 0
    left := 0
    right := len(height)-1
    for left < right {
        minNum := min(height[left], height[right])
        num := (right-left) * minNum
        maxNum = max(maxNum, num)
        if height[left] > height[right] {
            right --
        }else{
            left++
        }
    }

    return maxNum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a,b int)int{
    if a<b{
        return a
    }
    return b
}