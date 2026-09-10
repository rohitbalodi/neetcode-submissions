/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	output := []int{}
	queue := []*TreeNode{ root, nil}
	temp := []int{}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if front == nil {
			output = append(output, temp[len(temp)-1])
			temp = []int{}
			if len(queue)==0 {
				return output
			}
			queue = append(queue, nil)
		} else {
			temp = append(temp, front.Val)
			if front.Left != nil {
				queue = append(queue, front.Left)
			}

			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
	}

	return output
    
}