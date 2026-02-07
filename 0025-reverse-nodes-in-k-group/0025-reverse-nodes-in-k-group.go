/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	count := 0
	curr := head
	for curr != nil && count < k {
		curr = curr.Next
		count++
	}

	if count < k {
		return head
	}


	var prev, next *ListNode
	curr = head
	for i := 0; i < k; i++ {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head.Next = reverseKGroup(curr, k)
	return prev
}