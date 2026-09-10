/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	output := math.MinInt

	var dfs func(curr *TreeNode) int
	dfs = func(curr *TreeNode) int {
		if curr==nil {
			return 0
		}

		leftVal, rightVal := dfs(curr.Left), dfs(curr.Right)
		leftGain := max(leftVal, 0)
		rightGain := max(rightVal, 0)
		maxSoFar := leftGain + rightGain + curr.Val
		output = max(output, maxSoFar)
		return max(leftGain + curr.Val, rightGain + curr.Val, curr.Val)
	}
    
	dfs(root)
	return output
}
