/**
 * 有效的括号
 */
func isValid(s string) bool {
    arr:= make([]byte, len(s))
    n:=1
    arr[0] = s[0]
    for i:=1;i<len(s);i++{
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            arr[n]=s[i]
            n = n+1
        }else if n==0 {
            return false
        }else{
            if s[i] == ')' && arr[n-1] == '('{
                n=n-1
            }else if s[i] =='}' && arr[n-1] == '{' {
                n=n-1
            }else if s[i] == ']' && arr[n-1] == '['{
                n=n-1
            }else{
                return false
            }
        }
    }
    if n==0{
        return true
    }
    return false
}