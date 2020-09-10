/** 
 * 跳台阶
 */
func climbStairs(n int) int {
    
    //迭代方法超时了
//     if n == 1{
//         return 1
//     }
//     if n == 2 {
//         centerNum := climbStairs(n-1)
//         return centerNum+1
//     }
    
//     c1 := climbStairs(n-1)
//     c2 := climbStairs(n-2)
//     return c1+c2
    
    //解法二 
    a,b,c:=0,0,1
    for i:=1;i<=n;i++{
        a = b
        b = c
        c = a+b
    }
    
    return c
}