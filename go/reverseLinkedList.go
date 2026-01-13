// Note that this may have to modified to only reverse part of a list or smth
func reverseList(head *ListNode) *ListNode {

    var prev *ListNode
    curr := head

    for curr != nil {
        nextNode := curr.Next
        curr.Next = prev
        prev = curr
        curr = nextNode
    }

    return prev
}
