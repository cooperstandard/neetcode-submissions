/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
   balanced, _ := isBalancedHelper(root) 
   return balanced
}

func isBalancedHelper(root *TreeNode) (bool, int) {
    if root == nil {
        return true, 0
    }

    lBal, lHeight := isBalancedHelper(root.Left)
    rBal, rHeight := isBalancedHelper(root.Right)

    if !lBal || !rBal {
        return false, 0
    }

    diff := lHeight - rHeight

    if max(diff * -1, diff) > 1 {
        return false, 0
    }

    return true, max(lHeight, rHeight) + 1
}
