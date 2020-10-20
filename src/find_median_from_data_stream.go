/**
 *数据流的中位数
 *
 * 大小堆的应用 
 * 小顶堆存放数组较大的一部分数据，大顶堆存放数组较小的一部分数据，并且小顶堆的长度要比大顶堆的长度相等或者大于1
 * 中位数就是如果大小顶堆的长度一样，大小顶堆的堆顶的和的一半，如果小顶堆的长度比大顶堆的长度大，中位数是小顶堆的堆顶
 *
 * 
 */
type MedianFinder struct {
    minArr []int
    maxArr []int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    heap := MedianFinder {
        minArr: make([]int, 0),
        maxArr: make([]int, 0),
    }

    return heap
}


func (this *MedianFinder) AddNum(num int)  {
    if len(this.minArr) == 0 {
        this.minHeadUp(num)
        return
    }
    
    if num > this.minArr[0] {
        if len(this.minArr) + 1 - len(this.maxArr) > 1 {
            this.maxHeadUp(this.minArr[0])
            this.minHeadDown(num)
        }else {
            this.minHeadUp(num)
        }
    }else {
        if len(this.maxArr) + 1 > len(this.minArr) {
            if num < this.maxArr[0] {
                this.minHeadUp(this.maxArr[0])
                this.maxHeadDown(num)
            }else{
                this.minHeadUp(num)
            }
        }else {
            this.maxHeadUp(num)
        }
    }
}


func (this *MedianFinder) FindMedian() float64 {
    // fmt.Println(this.minArr, this.maxArr)
    if len(this.minArr) == len(this.maxArr) {
        return ((float64(this.minArr[0]) + float64(this.maxArr[0])) / 2)
    }else {
        return float64(this.minArr[0])
    }
}

func (this *MedianFinder) minHeadUp(val int) {
    this.minArr = append(this.minArr, val)
    lens := len(this.minArr) - 1
    for lens > 0 && this.minArr[lens] < this.minArr[(lens-1)/2] {
        swap(this.minArr, lens, (lens-1)/2)
        lens = (lens-1)/2
    }
    
}

func (this *MedianFinder) maxHeadUp(val int) {
    this.maxArr = append(this.maxArr, val)
    lens := len(this.maxArr) - 1
    for lens > 0 && this.maxArr[lens] > this.maxArr[(lens-1)/2] {
        swap(this.maxArr, lens, (lens-1)/2)
        lens = (lens-1)/2
    } 
    
}

func (this *MedianFinder) minHeadDown(val int) {
    index := 0
    l := 2 * index + 1
    
    minest := 0
    this.minArr[0] = val
    for l < len(this.minArr) {
        if l + 1< len(this.minArr) && this.minArr[l] > this.minArr[l+1] {
            minest = l + 1
        }else {
            minest = l
        }
        
        if this.minArr[index] < this.minArr[minest] {
            minest = index
        }
        
        if index == minest {
            break
        }
        
        swap(this.minArr, index, minest)
        index = minest
        l = 2 * index + 1
    }
}

func (this *MedianFinder) maxHeadDown(val int) {
    index := 0
    l := 2 * index + 1
    
    maxest := 0
    this.maxArr[0] = val
    for l < len(this.maxArr) {
        if l + 1 < len(this.maxArr) && this.maxArr[l] < this.maxArr[l+1] {
            maxest = l + 1
        }else {
            maxest = l
        }
        
        if this.maxArr[index] > this.maxArr[maxest] {
            maxest = index
        }
        
        if index == maxest {
            break
        }
        
        swap(this.maxArr, index, maxest)
        index = maxest
        l = 2 * index + 1
    }
}

func swap(arr []int, min, max int) {
    arr[min],arr[max] = arr[max],arr[min]
} 
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */