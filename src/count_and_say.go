/** 
 * 外观数列
 * 解题思路：用迭代，然后就是string的一些用法
 * 
 */
func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    if n == 2 {
        return "11"
    }
    
    lastStr := countAndSay(n-1)
    num := 1
    ch := lastStr[0]
    var buf bytes.Buffer
    for i:=1;i<len(lastStr);i++{
        if lastStr[i] == ch {
            num++
        }else{
            buf.WriteString(strconv.Itoa(num))
			buf.WriteString(string(ch))
            ch = lastStr[i]
            num = 1
        }
    }
    
    buf.WriteString(strconv.Itoa(num))
	buf.WriteString(string(ch))
    return buf.String()
}