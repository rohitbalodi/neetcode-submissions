/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	output := []int{}
	var dfs func(curr *TreeNode)
	dfs = func(curr *TreeNode) {
		if curr == nil {
			return
		}

		dfs(curr.Left)
		output = append(output, curr.Val)
		dfs(curr.Right)
	}

	dfs(root)
	return output[k-1]
    
}
