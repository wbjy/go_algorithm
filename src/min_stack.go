/** 
 *最小栈
 *
 *  先入后出的思想
 * 
 */

type MinStack struct {
    min int
    arr []int

}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        0,
        []int{},
    }
}


func (this *MinStack) Push(x int)  {
    if this.min > x || len(this.arr) == 0{
        this.min = x
    }
    this.arr = append(this.arr, x)
}


func (this *MinStack) Pop()  {
    if this.Top() == this.min {
		for i := 0; i < len(this.arr)-1; i++ {
			if i == 0 || this.arr[i] < this.min {
				this.min = this.arr[i]
			}
		}
	}
    this.arr = this.arr[:len(this.arr)-1]
    
}


func (this *MinStack) Top() int {
    return this.arr[len(this.arr)-1]
}


func (this *MinStack) GetMin() int {
    return this.min
}