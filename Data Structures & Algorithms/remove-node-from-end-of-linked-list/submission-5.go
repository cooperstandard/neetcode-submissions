/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ret := &ListNode{Next: head}
	p1, p2 := ret, ret

	for i := -1; i < n; i++ {
		p2 = p2.Next
	}

	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p1.Next = p1.Next.Next


	return ret.Next
}
