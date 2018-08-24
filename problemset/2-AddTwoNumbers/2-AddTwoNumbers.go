package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nextadd := 0
	var curL1 *ListNode = l1
	var curL2 *ListNode = l2
	head := new(ListNode)
	thisDigit := head
	for {
		thisval := curL1.Val + curL2.Val + nextadd
		thisDigit.Val = thisval % 10
		nextadd = thisval / 10
		if curL1.Next == nil && curL2.Next == nil {
			break
		}
		if curL1.Next != nil {
			curL1 = curL1.Next
		} else {
			curL1 = new(ListNode)
		}
		if curL2.Next != nil {
			curL2 = curL2.Next
		} else {
			curL2 = new(ListNode)
		}
		thisDigit.Next = new(ListNode)
		thisDigit = thisDigit.Next
	}
	if nextadd != 0 {
		thisDigit.Next = new(ListNode)
		thisDigit.Next.Val = nextadd
	}
	return head
}
