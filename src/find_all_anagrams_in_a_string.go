/**
 *找到字符串中所有字母异位词
 *
 * 解题思路：滑动窗口
 * 
 */
func findAnagrams(s string, p string) []int {
    //解法一：这个算是滑动窗口的一种感觉
//     if len(s) < len(p) {
//         return []int{}
//     }
//    res := make([]int, 0)
//    m := make(map[byte]int)
//    for _,v := range p {
//        m[byte(v)] ++
//    }

//    for i:= 0;i<len(p);i++{
//        m[s[i]] --
//    }

//    for i:=0;i<len(s)-len(p)+1;i++{
//        flag := 0
//        if i > 0{
//            m[s[i-1]] ++
//            m[s[i+len(p)-1]] --
//        }

//        for _,v := range m {
//            if v != 0{
//                flag = 1
//                break
//            }
//        }

//        if flag == 0{
//            res = append(res, i)
//        }
//    }
//     return res
    //解法二：滑动窗口
    need := make(map[byte]int)
    win := make(map[byte]int)
    for i:=0;i<len(p);i++{
        need[p[i]] ++
    }
    res := make([]int, 0)
    start, end := 0, 0
    l := 0
    for end < len(s){
        c:=s[end]
        end ++
       if _, ok := need[c]; ok {
            win[c] ++
            if win[c] == need[c] {
                l++
            }
        }
        for end - start >= len(p) {
            if l == len(need) {
                res = append(res, start)
            }
            d := s[start]
            start++
            if _, ok := win[d];ok{
                if win[d] == need[d] {
                    l--
                }
                win[d] --
            }
        }
    }
    return res


}