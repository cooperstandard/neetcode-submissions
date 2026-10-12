/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ret := &ListNode{}
	cursor := ret

	for list1 != nil || list2 != nil {

		if list1 == nil {
			cursor.Next = list2
			break
		}

		if list2 == nil {
			cursor.Next = list1
			break
		}


		if list1.Val < list2.Val {
			cursor.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			cursor.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		cursor = cursor.Next




	}

	return ret.Next
    
}
