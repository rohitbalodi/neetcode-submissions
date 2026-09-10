/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	var dfs func(curr *TreeNode) (*TreeNode)
	dfs = func(curr *TreeNode) (*TreeNode) {
		if curr == nil {
			return nil
		}

		leftNode, rightNode := dfs(curr.Left), dfs(curr.Right)
		curr.Left = rightNode
		curr.Right = leftNode
		return curr
	}

	dfs(root)
	return root
    
}
