/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
   fast := head
   slow := head
   for fast != nil && slow != nil {
     if fast.Next == nil {
        return false
     }
     fast = fast.Next
     if slow == fast {
        return true
     }
     slow = slow.Next
     if fast.Next == nil {
        return false
     }
     fast = fast.Next
   }

    return false

}
