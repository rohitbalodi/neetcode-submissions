/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{ root, nil}
	for len(queue)>0 {
		front := queue[0]
		queue = queue[1:]
		// since its a nil, need not set next here
		if front == nil {
			// last nil
			if len(queue)==0 {
				return root
			}
			// maybe change of level
			queue = append(queue, nil)
		} else {
			// so let's check if these are full nodes
			if front.Left != nil && front.Right != nil {
				queue = append(queue, front.Left)
				queue = append(queue, front.Right)
			}
			// else do nothing since they are leaf nodes
			front.Next = queue[0]
		}

	}

	return root
	
}
