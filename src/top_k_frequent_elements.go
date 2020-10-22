/**
 * 前 K 个高频元素
 * 
 */


// 解法一  运用小顶堆  非常之垃圾的写法
// func topKFrequent(nums []int, k int) []int {
//     m := make(map[int]int)
//     for _,v := range nums {
//         m[v] ++
//     }
    
//     minArr := []int{}
//     for _,value := range m {
//         if len(minArr) < k {
//             minArr = minHeadUp(minArr, value)
//         }else{
//             if value <= minArr[0] {
//                 continue
//             }else{
//                 minArr = minHeadDown(minArr, value)
//             }
//         }
//     }
//     fmt.Println(minArr)
//     res := []int{}
    
//     for _ , v := range minArr {
//        for key, mv := range m {
//             if v == mv {
//                 res = append(res, key)
//                 m[key] = 0
//             }
//         }
//     }
    
//     return res
// }

// func minHeadUp(minArr []int, val int) []int {
//     if len(minArr) == 0 {
//         minArr = append(minArr, val)
//         return minArr
//     }
//     minArr = append(minArr, val)
//     i := len(minArr) - 1
//     for i > 0 && minArr[i] < minArr[(i-1)/2] {
//         minArr = swap(minArr, i, (i-1)/2)
//         i = (i-1)/2
//     }
    
//     return minArr
// }

// func minHeadDown(minArr []int, val int) []int {
//     minArr[0] = val
//     index := 0
//     i := index * 2 + 1
//     minest := 0
//     for i < len(minArr) {
        
//         if i + 1 < len(minArr) && minArr[i] > minArr[i + 1] {
//             minest = i + 1
//         }else {
//             minest = i
//         }
        
//         if minArr[index] < minArr[minest]{
//             minest = index
//         }
        
//         if index == minest {
//             break
//         }
        
//         minArr = swap(minArr, minest, index)
//         index = minest
//         i = 2*index + 1
//     }
    
//     return minArr
// }

// func swap(arr []int, i, j int) []int {
//     arr[i], arr[j] = arr[j], arr[i]
//     return arr
// }

// 解法二：运用快排思想
func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for _, v := range nums {
        m[v] ++
    }
    
    values := [][]int{}
    for key,value := range m {
        values = append(values, []int{key,value})
    }
    
    res := make([]int, k)
    
    quickSort(values, 0, len(values)-1, 0, k, res)
    
    return res
}

func quickSort(values [][]int, start, end, count, k int, res []int) {
     rand.Seed(time.Now().UnixNano())
    picked := rand.Int() % (end - start + 1) + start
    values[start],values[picked] = values[picked],values[start]
    pivot := values[start][1]
    index := start
    for i := start+1; i <= end; i++{
        if values[i][1] >= pivot {
            values[index+1],values[i] = values[i],values[index+1]
            index ++
        }
    }
    
    values[index],values[start] = values[start],values[index]
    
    if index - start >= k {
        quickSort(values, start, index-1, count, k, res)
    }else {
        for i := start; i <= index; i++{
            res[count] = values[i][0]
            count ++
        }
        if k > (index - start + 1) {
            quickSort(values, index+1, end, count, k - (index - start + 1), res)
        }
    }
}
