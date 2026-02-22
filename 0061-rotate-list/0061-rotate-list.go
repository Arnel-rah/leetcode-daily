/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }

    length := 1
    tail := head
    for tail.Next != nil {
        tail = tail.Next
        length++
    }

    k %= length
    if k == 0 {
        return head
    }

    cur := head
    for i := 1; i < length-k; i++ {
        cur = cur.Next
    }

    newHead := cur.Next
    cur.Next = nil
    tail.Next = head

    return newHead
}