package _114_二叉树展开为链表
// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/


 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }


 func main(){
	 var left *TreeNode
	 var right *TreeNode
	 root := &TreeNode{10,left,right}
	 flatten(root)
 }

func flatten(root *TreeNode)  {
    list := preorderTraversal(root)
    for i := 1; i < len(list); i++ {
        prev, curr := list[i-1], list[i]
        prev.Left, prev.Right = nil, curr
    }
}

func preorderTraversal(root *TreeNode) []*TreeNode {
    list := []*TreeNode{}
    if root != nil {
        list = append(list, root)
        list = append(list, preorderTraversal(root.Left)...)
        list = append(list, preorderTraversal(root.Right)...)
    }
    return list
}

// @lc code=end

