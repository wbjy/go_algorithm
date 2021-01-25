/**
 *全排列
 *
 * 解题思路：这种需要列出所有结果的一般都可以用回溯来解决，枚举所有解法，通过剪枝来优化
 *
 * 
 */

var res [][]int
func permute(nums []int) [][]int {
    res = [][]int{}
    m := make(map[int] bool)
    pathNums := []int{}
    dfs(nums, pathNums, m)
    return res
}

func dfs(nums []int, pathNums []int, m map[int]bool){
    if len(nums) == len(pathNums) {
        temp := make([]int, len(nums))
        copy(temp, pathNums)    //因为pathNums是引用类型的，如果这里不copy一下的话之后修改会影响到最后结果
        res = append(res, temp)
        return
    }
    for i:=0;i<len(nums);i++{
        if !m[nums[i]] { //没被选过
            pathNums = append(pathNums, nums[i])
            m[nums[i]] = true                       //选择状态
            dfs(nums, pathNums, m)                     //
            pathNums = pathNums[:len(pathNums)-1]   //倒退选择
            m[nums[i]] = false                      //撤销选择状态
        }
    }
}