/** 
 *字符串转化成数字
 *
 * 
 */

func myAtoi(str string) int {
    str = strings.TrimSpace(str)
    a := 1
    num := 0
    for k,v := range str{
        if v >= '0' && v <= '9' {
            num = num*10+int(v-'0')
        }else if v == '-' && k == 0{
            a = -1
        }else if v == '+' && k == 0{
            a = 1
        }else{
            break
        }
        
        if num > math.MaxInt32 {
            if a == -1 {
                return math.MinInt32
            }
            return math.MaxInt32
        }
    }
    return num*a
}