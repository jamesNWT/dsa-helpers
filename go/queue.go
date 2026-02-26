
type ListNode struct {
    Val  int
    Next *ListNode
	Prev *ListNode
}

type Queue struct {
	Left *ListNode
	Right *ListNode
}

func newQueue() *Queue {
	ret := &Queue{Left: &ListNode{}, Right: &ListNode{}}
	ret.Right.Prev = ret.Left
	ret.Left.Next = ret.Right
	return ret
}

func (q *Queue) enqueue(x int) {
	newNode := &ListNode{Val: x, Next: q.Right, Prev: q.Right.Prev}
	q.Right.Prev.Next = newNode
	q.Right.Prev = newNode
}


func (q *Queue) dequeue() int {
	ret := q.Left.Next.Val
	q.Left.Next = q.Left.Next.Next
	q.Left.Next.Prev = q.Left
	return ret
}

func (q *Queue) toArray() []int {
	var ret []int
	dummy := q.Left.Next
	for dummy != q.Right {
		ret = append(ret, dummy.Val)
		dummy = dummy.Next
	}
	return ret
}
