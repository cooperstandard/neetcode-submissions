/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	levels := [][]int{}
	 q := []*TreeNode{root}


	for len(q) != 0 {
		newChildren := []*TreeNode{}
		level := []int{}
		for len(q) != 0 {
			n := q[0]
			q = q[1:]
			if n == nil {
				continue
			}
			level = append(level, n.Val)
			newChildren = append(newChildren, n.Left, n.Right)
		}
		levels = append(levels, level)
		for _, newChild := range newChildren {
			q = append(q,newChild)
		}
	}
	return levels[:len(levels)-1]

    
}

