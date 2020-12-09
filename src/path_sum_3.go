/**
 *路径总和 III
 *
 * 解题思路：DFS+回溯+前缀
 * 
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var m map[int]int
func pathSum(root *TreeNode, sum int) int {
    m = make(map[int]int)
    m[0] = 1
    return dfs(root, sum, 0)
}

func dfs(root *TreeNode, sum, numSum int) int{
    if root == nil {
        return 0
    }
    var res int
    numSum = numSum + root.Val
    res =res + m[numSum-sum]
    
    m[numSum] = m[numSum] + 1
    res =res + dfs(root.Left, sum, numSum)
    res = res + dfs(root.Right, sum, numSum)
    m[numSum] = m[numSum] - 1
    return res
}