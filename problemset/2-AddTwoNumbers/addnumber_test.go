package addtwonumbers

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	Number1 := new(ListNode)
	Number1.Val = 2
	Number1.Next = new(ListNode)
	Number1.Next.Val = 4
	Number1.Next.Next = new(ListNode)
	Number1.Next.Next.Val = 3

	Number2 := new(ListNode)
	Number2.Val = 5
	Number2.Next = new(ListNode)
	Number2.Next.Val = 6
	Number2.Next.Next = new(ListNode)
	Number2.Next.Next.Val = 4

	for NumberSum := addTwoNumbers(Number1, Number2); NumberSum.Next != nil; NumberSum = NumberSum.Next {
		fmt.Println(NumberSum.Val)
	}
}
