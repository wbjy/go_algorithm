func longestCommonPrefix(strs []string) string {
    if len(strs) == 0{
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }
    var buf bytes.Buffer
    i := 0
    j := 0
    var cenByte byte
    for {
        s := strs[i]
        if len(s)-1 < j{
            return buf.String()
        }
        if s[j] == cenByte || i == 0{
            if i == 0{
                cenByte = s[j]
                i++
            }else if i == len(strs) - 1{
                buf.WriteString(string(cenByte))
                i = 0
                j++
            }else {
                i ++
            }
        }else {
            return buf.String()
        }
    }
}