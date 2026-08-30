/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		// store the next
		next := curr.Next
		// change current next to previous
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
	
    
}
