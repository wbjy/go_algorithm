/**
 *二叉树的序列化与反序列化
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

type Codec struct {
    l []string
}

func Constructor() Codec {
    return Codec{}
}

func reserialize(root *TreeNode, str string) string {
    if root == nil {
        str += "null,"
    }else{
        str += strconv.Itoa(root.Val) + ","
        str = reserialize(root.Left, str)
        str = reserialize(root.Right, str)
    }
    return str
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    return reserialize(root, "")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    l := strings.Split(data, ",")
    for i:= 0; i< len(l); i++{
        if l[i] != ""{
            this.l = append(this.l, l[i])
        }
    }
    return this.redeserialize()

}

func (this *Codec) redeserialize() *TreeNode {
    if this.l[0] == "null" {
        this.l = this.l[1:]
        return nil
    }
    val, _ := strconv.Atoi(this.l[0])
    root := &TreeNode{Val:val}
    this.l = this.l[1:]
    root.Left = this.redeserialize()
    root.Right = this.redeserialize()
    return root
}