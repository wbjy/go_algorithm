/*
    解法一：快速排序

func findKthLargest(nums []int, k int) int {
    idx := quickSort(0, len(nums) - 1, len(nums) - k, nums)
    return nums[idx]
}

func quickSort(l, r, pos int, nums []int) int {
	pIndex := partition(nums, l, r)
    if pIndex == pos {
        return pIndex
    }else if pIndex < pos {
		return quickSort(pIndex+1, r, pos, nums)
    }else{
        return quickSort(l, pIndex-1, pos, nums)
    }
}

func partition(arr []int, l, r int) int {
    s := l
    cen_val := arr[l]
    for l < r {
        
        for ;l < r && arr[r] >= cen_val;{
            r--
        }
        for ;l < r && arr[l] <= cen_val;{
            l++
        }
        
        
        if l < r{
            swap(arr, l, r)
        }
    }
    
    arr[s] = arr[l]
    arr[l] = cen_val
	return l
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
*/

//解法二  小顶堆+堆排序
func findKthLargest(nums []int, k int) int {
    nums = headSort(nums)
    return nums[k-1]
}

func headSort(nums []int) []int{
    lens := len(nums)
    buildHead(nums, lens)
    
    for i := lens-1; i >= 0; i--{
        swap(nums, 0, i)
        head(nums, 0, i)
    }
    
    return nums
}

func buildHead(nums []int, lens int){
    for i:= lens/2; i>=0; i-- {
        head(nums, i, lens)
    }
}

func head(nums []int, i, lens int) {
    min := i
    l := i*2+1
    r := i*2+2
    if l < lens && nums[l] < nums[min] {
        min = l
    }
    
    if r < lens && nums[r] < nums[min] {
        min = r
    }
    
    if min != i {
        swap(nums, i, min)
        head(nums, min, lens)
    }
}


func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
