/** 
 *验证回文串
 *
 *给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 * 
 */
func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    j := len(s)-1
    i := 0
    for i<j {
        flag1 := getChar(s[i])
        flag2 := getChar(s[j])
        if flag1 && flag2 {
            if s[i] == s[j] {
                i ++
                j --
            }else{
                return false
            }
        }else if ! flag1 && flag2{
            i++
        }else if flag1 && ! flag2{
            j--
        }else{
            i ++
            j --
        }
        
    }
    
    return true
    
}

func getChar(s byte) bool {
    return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9')
}