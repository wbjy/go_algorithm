/**
 * 多数元素
 */

func majorityElement(nums []int) int {
    
    //解法一
//     m := make(map[int]int)
//     for _,v := range nums{
//         m[v] ++
//     }
//     maxV := 0
//     maxK := 0
//     for k,v1 := range m{
//         if maxV < v1 {
//             maxV = v1
//             maxK = k
//         }
//     }
    
//     return maxK
    
    //解法二
    centerNum := nums[0]
    count := 1
    for _,v := range nums[1:]{
        if count == 0{
            centerNum = v
        }
        if centerNum == v{
            count++
        }else{
            count--
        }
    }
    
    return centerNum
    
}