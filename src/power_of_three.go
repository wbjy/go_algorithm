/**
 * 3的幂
 */
func isPowerOfThree(n int) bool {
    //解法一
	// if n < 1 {
	// 	return false
	// }
	// s := strconv.FormatInt(int64(n), 3)
	// return s[0:1] == "1" && strings.Count(s, "0") == len(s)-1
    
    //解法二
    if n==0{
        return false
    }
    if n != 1{
        for n%3==0{
            n=n/3
        }
    }
    return n==1
}