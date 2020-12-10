/**
 *字母异位词分组
 *
 * 解题思路：大体思路就是用同一个标示来代表异位字符串，可以是数组，排序后的string等，然后用哈希表来对应
 * 
 */

// 解法一：计数
// func groupAnagrams(strs []string) [][]string {
//     mStr := make(map[[26]int][]string)
//     for _, str := range strs {
//         key := getKey(str)
//         s, ok := mStr[key]
//         if !ok{
//             s = make([]string, 0)
//         }
//         s = append(s, str)
//         mStr[key] = s
//     }
//     var res [][]string
//     for _, arr := range mStr {
//         res = append(res, arr)
//     }
//     return res
// }

// func getKey(str string) [26]int{
//     res := [26]int{}
//     for _, v := range str {
//         res[int(v-'a')] ++
//     }

//     return res
// }

//解法二：排序
func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    for _, str := range strs {
        key := getKey(str)
        s, ok := m[key]
        if !ok{
            s = make([]string, 0)
        }
        s = append(s, str)
        m[key] = s
    }
    var res [][]string
    for _, arr := range m {
        res = append(res, arr)
    }
    return res
}

func getKey(str string) string{
    bytes := []byte(str)
    sort.Slice(bytes, func(i, j int) bool{return bytes[i] < bytes[j]})
    return string(bytes)
}