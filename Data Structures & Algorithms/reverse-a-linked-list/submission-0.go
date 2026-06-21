/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	current := head
	next := current.Next
	for next != nil {
		current.Next = prev
		prev = current
		current = next
		next = current.Next
	}
	current.Next = prev

	return current
}
