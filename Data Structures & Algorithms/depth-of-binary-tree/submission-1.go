/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	var dfs func(curr *TreeNode) (int)
	dfs = func(curr *TreeNode) (int) {
		if curr==nil {
			return 0
		}

		leftVal, rightVal := dfs(curr.Left), dfs(curr.Right)
		return max(leftVal, rightVal)+1
	}

	return dfs(root)

    
}
