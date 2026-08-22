/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
     if root == nil || (root.Left == nil && root.Right == nil) {
  } else if root.Left == nil {
    root.Left = invertTree(root.Right)
    root.Right = nil
  } else if root.Right == nil {
    root.Right = invertTree(root.Left)
    root.Left = nil
  } else {
    temp := invertTree(root.Right)
    root.Right = invertTree(root.Left)
    root.Left = temp
  }

  return root 
}
