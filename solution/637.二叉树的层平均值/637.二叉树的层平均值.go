/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
 *
 * https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/description/
 *
 * algorithms
 * Easy (68.16%)
 * Likes:    176
 * Dislikes: 0
 * Total Accepted:    39.9K
 * Total Submissions: 58.5K
 * Testcase Example:  '[3,9,20,15,7]'
 *
 * 给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 输出：[3, 14.5, 11]
 * 解释：
 * 第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 节点值的范围在32位有符号整数范围内。
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 package _637_二叉树的层平均值

 type TreeNode struct{
	 Val int
	 Left *TreeNode
	 Right *TreeNode
 }
 
func averageOfLevels(root *TreeNode) []float64 {
    if root == nil {
        return []float64{0}
	}
	list, average := []*TreeNode{root}, make([]float64, 0)
    bfs(list, &average)
    return average
}
// @lc code=end

func bfs(list []*TreeNode, average *[]float64){
    var sum, num float64 = 0, 0
    newList := make([]*TreeNode, 0)
    for _,v := range list{
        sum += float64(v.Val)
        num++
        if v.Left != nil {
            newList = append(newList, v.Left)
        }
        if v.Right != nil {
            newList = append(newList, v.Right)
        }
    }
    if num>0{
        *average = append(*average, sum/num)
    }
    if len(newList)>0{
        bfs(newList, average)
    }    
}