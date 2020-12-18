/**
 *寻找两个正序数组的中位数
 *
 * 解题思路：二分法
 * 
 */

// 解法一：暴力解法
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     length := len(nums1) + len(nums2)
//     nums := make([]int, length)
//     start1, start2 := 0,0
//     var mid float64
//     for i:=0;i<length;i++{
//         if start1 < len(nums1) && start2 < len(nums2) {
//             if nums1[start1] < nums2[start2] {
//                 nums[i] = nums1[start1]
//                 start1++
//             }else{
//                 nums[i] = nums2[start2]
//                 start2++
//             }
//         }else if start1 < len(nums1) {
//             nums[i] = nums1[start1]
//             start1 ++
//         }else if start2 < len(nums2) {
//             nums[i] = nums2[start2]
//             start2 ++
//         }

//         if length % 2 == 0 {
//             if i >= length/2 {
//                 mid = float64(nums[i-1]+nums[i])/2.0
//                 return mid
//             }
//         }else {
//             if i >= length/2 {
//                 mid = float64(nums[i])
//                 return mid
//             }
//         }
//     }

//     return mid
// }

//二分法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    length := len(nums1) + len(nums2)
    k := length/2
    if length % 2 == 0 {
        return (float64(getKth(nums1, nums2, k))+float64(getKth(nums1, nums2, k+1)))/2.0
    }else{
        return float64(getKth(nums1, nums2, k + 1))
    }
}

func getKth(nums1 []int, nums2 []int, k int) int {
    index1, index2 := 0, 0
    for {
        if index1 == len(nums1) {
            return nums2[index2 + k - 1]
        }
        if index2 == len(nums2) {
            return nums1[index1 + k - 1]
        }

        if k == 1 {
            return min(nums1[index1], nums2[index2])
        }
        half := k/2
        newIndex1 := min(index1+half, len(nums1)) - 1
        newIndex2 := min(index2+half, len(nums2)) - 1
        pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
        if pivot1 <= pivot2 {
            k -= (newIndex1 - index1 + 1)
            index1 = newIndex1 + 1
        }else{
            k -= (newIndex2 - index2 + 1)
            index2 = newIndex2 + 1
        }
    }
    return 0
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}