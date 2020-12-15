/**
 * 打家劫舍 III
 *
 * 解题思路：动态规划   主要是找出状态转移方程，也就是第27，28行代码描述
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
func rob(root *TreeNode) int {
    v := dfs(root)
    return max(v[0], v[1])
}

func dfs(root *TreeNode) []int {
    if root ==nil {
        return []int{0, 0}
    }

    l, r := dfs(root.Left), dfs(root.Right)
    sel := root.Val + l[1] + r[1]
    noSel := max(l[0], l[1]) + max(r[0], r[1])
    return []int{sel, noSel}
}

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}