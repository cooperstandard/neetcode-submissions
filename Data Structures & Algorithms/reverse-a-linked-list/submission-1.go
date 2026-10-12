/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var reverse func(current, prev *ListNode) *ListNode 
	reverse = func(current, prev *ListNode) *ListNode {
		if current == nil {
			return nil
		}

		if current.Next == nil {
			current.Next = prev
			return current
		}

		next := current.Next
		current.Next = prev
		return reverse(next, current)

	}

	return reverse(head, nil)
}
