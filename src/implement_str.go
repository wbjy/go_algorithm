/** 
 *实现 strStr()
 * 
 */

func strStr(haystack string, needle string) int {
    if needle == ""{
        return 0
    }else if haystack == ""{
        return -1
    }
    i := 0
    j := 0
    oldi := -1
    for i < len(haystack) {
        if haystack[i] == needle[j]{
            if oldi == -1 {
                oldi = i
            }
            
            i ++
            j ++
        }else {
            if oldi >= 0 {
                i = oldi
                oldi = -1
            }
            i ++
            j = 0
            if i > len(haystack) - len(needle) {
                return -1
            }
            
        }
        if j>=len(needle){
            return i-j
        }
    }
    
    return -1

}