/** 
 * 基本计算器 II
 */
func calculate(s string) int {
    signs := make([]int, 0)
    stack := make([]int, 0)
    num := 0
    for i:=0;i<len(s);i++{
        switch s[i] {
            case '1','2','3','4','5','6','7','8','9','0':
                num = 0
                for i<len(s)&&s[i] >= '0' && s[i] <= '9'{
                    num = num*10+int(s[i]-'0')
                    i ++
                }
            if len(signs) >0 && signs[len(signs)-1] > 2 {
                if signs[len(signs)-1] == 3 {
                    stack[len(stack)-1] = stack[len(stack)-1] * num
                }else {
                    stack[len(stack)-1] = stack[len(stack)-1] / num
                }
                signs = signs[:len(signs)-1]
            }else{
                stack = append(stack, num)
            }
            i --
            case '+':
            signs = append(signs, 1)
            case '-':
            signs = append(signs, -1)
            case '*':
            signs = append(signs, 3)
            case '/':
            signs = append(signs, 4)
            default:
        }
    }
    
    for len(signs) > 0{
        stack[1] = stack[0] + signs[0]*stack[1]
        signs = signs[1:]
        stack = stack[1:]
    }
    
    return stack[0]

}