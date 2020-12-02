/**
 *从前序与中序遍历序列构造二叉树
 *
 * 解题思路：递归
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
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0{
        return nil
    }
    var root = new(TreeNode)
    root.Val = preorder[0]
    index := 0
    for i:= 0;i<len(inorder);i++{
        if inorder[i] == preorder[0] {
            index = i
            break
        }
    }

    l_preorder := preorder[1:index+1]
    l_inorder := inorder[0:index]
    root.Left = buildTree(l_preorder, l_inorder)

    r_preorder := preorder[index+1:]
    r_inorder := inorder[index+1:]
    root.Right = buildTree(r_preorder, r_inorder)
    return root
}