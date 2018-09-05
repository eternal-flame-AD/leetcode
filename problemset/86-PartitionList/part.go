package part

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var head1 *ListNode
	var tail1 *ListNode
	var head2 *ListNode
	var tail2 *ListNode
	for head != nil {
		newnode := &ListNode{
			head.Val,
			nil,
		}
		if head.Val < x {
			if head1 != nil {
				tail1.Next = newnode
				tail1 = newnode
			} else {
				head1 = newnode
				tail1 = newnode
			}
		} else {
			if head2 != nil {
				tail2.Next = newnode
				tail2 = newnode
			} else {
				head2 = newnode
				tail2 = newnode
			}
		}
		head = head.Next
	}
	if head1 != nil {
		tail1.Next = head2
		return head1
	} else {
		return head2
	}
}
