/**
 *翻转二叉树
 *
 * 解题思路：递归大法
 * 
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root ==nil {
        return root
    }
    root.Left,root.Right = root.Right,root.Left
    invertTree(root.Left)
    invertTree(root.Right)
    return root
}