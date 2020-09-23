/** 
 * 打乱数组
 */

type Solution struct {
    arr []int
}


func Constructor(nums []int) Solution {
    return Solution{nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.arr
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    n := len(this.arr)
    res := make([]int, n)
    copy(res, this.arr)
    for i:=n-1;i>=0;i--{
        a := rand.Intn(i+1)
        res[i],res[a] = res[a],res[i]
    }
    
    return res
}
