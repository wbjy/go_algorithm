/**
 * 有序矩阵中第K小的元素
 *
 * 解题思路：解法一：将二维数组变成一位数组然后排序，这种解法没有运用的到二维有序数组的特性，时间复杂度较高
 *           解法二：运用二分法，取中位数，将数组分为上半区和下半区，这种解法运用到了二维有序数组的全部特性
 */

/** 解法一： 大顶堆 维护一个长度为K的大顶堆，最后堆顶的值就是第K小的值*/
// func kthSmallest(matrix [][]int, k int) int {
//     maxArr := []int{}
//     for i:=0;i<len(matrix);i++{
//         for j:=0;j<len(matrix[0]);j++{
//             if len(maxArr) < k {
//                 maxArr = maxHeadUp(maxArr, matrix[i][j])
//             }else{
//                 if maxArr[0] < matrix[i][j] {
//                     continue
//                 }else {
//                     maxArr = maxHeadDown(maxArr, matrix[i][j])
//                 }
//             }
//         }
//     }
//     return maxArr[0]
// }

// func maxHeadUp(maxArr []int, val int) []int {
//     if len(maxArr) == 0 {
//         maxArr = append(maxArr, val)
//         return maxArr
//     }
//     maxArr = append(maxArr, val)
//     i := len(maxArr) - 1
//     for i > 0 && maxArr[i] > maxArr[(i-1)/2] {
//         maxArr = swap(maxArr, i, (i-1)/2)
//         i = (i-1)/2
//     }
    
//     return maxArr
// }

// func maxHeadDown(maxArr []int, val int) []int {
//     maxArr[0] = val
//     index := 0
//     i := index * 2 + 1
//     maxest := 0
//     for i < len(maxArr) {
        
//         if i + 1 < len(maxArr) && maxArr[i] < maxArr[i + 1] {
//             maxest = i + 1
//         }else {
//             maxest = i
//         }
        
//         if maxArr[index] > maxArr[maxest]{
//             maxest = index
//         }
        
//         if index == maxest {
//             break
//         }
        
//         maxArr = swap(maxArr, maxest, index)
//         index = maxest
//         i = 2*index + 1
//     }
    
//     return maxArr
// }

// func swap(arr []int, i, j int) []int {
//     arr[i], arr[j] = arr[j], arr[i]
//     return arr
// }


/** 解法二 二分法 取中位数，然后切割上半区和下半区，最后在判断半区长度和K的大小
            check方法里面还可以优化一下
*/
func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    l, r := matrix[0][0],matrix[n-1][n-1]
    for l < r{
        mid := l + (r - l)/2
        if check(matrix, n, mid, k) {
            r = mid
        }else {
            l = mid + 1
        }
    }
    return l
}

func check(matrix [][]int, n, mid, k int) bool {
    num := 0
    for i := n - 1;i >= 0; i--{
        for j := 0;j < n; j ++ {
            if matrix[i][j] <= mid {
                num ++ 
            }
        }
    }
    return num >= k
}