/** 
 * 有效的字母异位词
 *
 * 两个字符串有相同的字母个数，单个字母的数量也要一样
 */
func isAnagram(s string, t string) bool {
    m := make(map[rune]int)
    for _,v := range s {
        m[v] ++
    }
    for _, v := range t {
        vs,ok := m[v]
        if !ok {
            return false
        }
        
        if vs == 1{
            delete(m, v)
        }else{
            m[v] --
        }
    }
    return len(m) == 0
}