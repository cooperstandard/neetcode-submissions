/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	// strategy, traverse the tree twice, first time to find furthest left and second to find furthest right
	// do so recursively with a depth first search approach
	// return findMostLeft(root) + findMostRight(root)
	  res := 0

    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }

        left := dfs(root.Left)
        right := dfs(root.Right)
        res = max(res, left + right)

        return 1 + max(left, right)
    }

    dfs(root)
    return res
    
}


func getMaxDiameter(root *TreeNode) (diameter, depth int ) {
	if root == nil {
		return 0, 0
	}

	left, dLeft := getMaxDiameter(root.Left)

	right, dRight := getMaxDiameter(root.Right)

	depth = max(right, left) + 1
	diameter = max(dRight, dLeft, 1 + left + right)
	return


}



func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

