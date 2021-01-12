/** 
 *
 *  二叉树的中序遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    arr1 := inorderTraversal(root.Left)
    arr1 = append(arr1, root.Val)
    arr2 := inorderTraversal(root.Right)
    arr1 = append(arr1, arr2...)
    return arr1
}