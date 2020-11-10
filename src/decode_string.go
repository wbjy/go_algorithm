/**
 *字符串解码
 *
 * 解题思路：用数组模拟栈的作用
 * 
 */
func decodeString(s string) string {
    charStack := []byte{}
    countStack := []int{}
    indexStack := []int{}
    count := 0
    
    for i:=0;i<len(s);i++{
        if s[i] >= '0' && s[i] <= '9' {
            count = count*10 + int(s[i] - '0')
        }else if s[i] == '[' {
            countStack = append(countStack, count)
            count = 0
            indexStack = append(indexStack, len(charStack))
        }else if s[i] == ']'{
            c := countStack[len(countStack)-1]
            countStack = countStack[:len(countStack) - 1]
            index := indexStack[len(indexStack)-1]
            indexStack = indexStack[:len(indexStack)-1]
            
            str := string(charStack[index:])
            charStack = charStack[:index]
            for ;c>0; c-- {
                charStack = append(charStack, []byte(str)...)
            }
        }else {
            charStack = append(charStack, s[i])
        }
    }
    
    return string(charStack)
    
}