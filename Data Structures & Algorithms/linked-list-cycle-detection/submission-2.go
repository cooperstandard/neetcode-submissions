/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head.Next, head

	for fast != nil && slow != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next

		slow = slow.Next
	}
	return false
    
}
