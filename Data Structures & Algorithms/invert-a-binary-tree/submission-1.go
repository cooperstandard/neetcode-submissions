/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	return swap(root)
}

func swap(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	swap(root.Left)
	swap(root.Right)

	return root
}