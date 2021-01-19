/**
 *
 *二叉树的最近公共祖先
 * 
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
   if root == nil {
       return nil
   }
   if root.Val == p.Val || root.Val == q.Val {
       return root
   }
   lNode := lowestCommonAncestor(root.Left, p, q)
   rNode := lowestCommonAncestor(root.Right, p, q)
   if lNode != nil && rNode != nil {
       return root
   }
   if lNode == nil {
       return rNode
   }
   return lNode
}