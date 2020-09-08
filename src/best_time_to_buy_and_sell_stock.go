/**
 * 买卖股票的最佳时机
 *
 * 解题思路：找出最小值和最小值后面的最大差值，并和前面作比较
 */
func maxProfit(prices []int) int {
   //解法一：效率极低 
//     a := 0
//     i := 1
//     j := 0
//     for j < len(prices)-1 {
//         if a < prices[i] - prices[j]{
//             a = prices[i] - prices[j]
//         }
//         if i == len(prices)-1{
//             j ++ 
//             i = j+1
//         }else{
//             i++
//         }
//     }
//     return a
    
    //解法二
    if len(prices)== 0 || len(prices)==1{
        return 0
    }
    min := prices[0]
    max := 0
    
    for i:=1;i<len(prices);i++{
        if prices[i] - min > max{
            max = prices[i] - min
        }
        if prices[i] < min{
            min = prices[i]
        }
    }
    return max
}