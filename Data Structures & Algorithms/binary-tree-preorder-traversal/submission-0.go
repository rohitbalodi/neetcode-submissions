/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    output := []int{}

	var dfs func(curr *TreeNode)
	dfs = func(curr *TreeNode) {
		if curr==nil {
			return
		}

		output = append(output, curr.Val)
		dfs(curr.Left)
		dfs(curr.Right)
	}

	dfs(root)
	return output
}
