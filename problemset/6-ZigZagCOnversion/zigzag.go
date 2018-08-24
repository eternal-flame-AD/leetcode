package zigzag

import (
	"bytes"
)

type rowChainNode struct {
	Val  byte
	Next *rowChainNode
}

type rowChain struct {
	Head *rowChainNode
	Tail *rowChainNode
}

func (c *rowChain) String() string {
	buff := bytes.NewBufferString("")
	curr := c.Head
	for curr != nil {
		buff.WriteByte(curr.Val)
		curr = curr.Next
	}
	return buff.String()
}

func (c *rowChain) Append(char byte) {
	newnode := new(rowChainNode)
	newnode.Val = char
	if c.Head == nil {
		c.Head = newnode
		c.Tail = newnode
	} else {
		c.Tail.Next = newnode
		c.Tail = newnode
	}
}

func convert(s string, numRows int) string {
	rows := make([]rowChain, numRows)
	currrow := 0
	currdelta := 1
	for index := range s {
		rows[currrow].Append(s[index])
		switch {
		case numRows == 1:
			break
		case currrow == 0 && currdelta == -1:
			currrow = 1
			currdelta = 1
		case currrow == numRows-1 && currdelta == 1:
			currrow = numRows - 2
			currdelta = -1
		default:
			currrow += currdelta
		}
	}
	res := ""
	for i := 0; i < numRows; i++ {
		res += rows[i].String()
	}
	return res
}
