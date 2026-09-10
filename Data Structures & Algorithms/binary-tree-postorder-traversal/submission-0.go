/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	output := []int{}

	var dfs func(curr *TreeNode)
	dfs = func(curr *TreeNode) {
		if curr==nil {
			return
		}

		dfs(curr.Left)
		dfs(curr.Right)
		output = append(output, curr.Val)
	}

	dfs(root)
	return output
    
}
