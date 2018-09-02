package ugly

/*
type TreeNode1 struct {
	number     int
	startindex int
}

func (c *TreeNode1) extend(primes []int) []*TreeNode1 {
	childs := make([]*TreeNode1, len(primes)-c.startindex)
	for index, val := range primes[c.startindex:] {
		childs[index] = &TreeNode1{c.number * val, index + c.startindex}
	}
	return childs
}

type QueueChain struct {
	head *QueueChainNode
}

func (c *QueueChain) insert(nodes []*TreeNode1, maxlength int) {
	cur := c.head
	count := 0
	for _, node := range nodes {
		if c.head == nil {
			c.head = &QueueChainNode{node, nil}
			cur = c.head
			continue
		}
		if node.number < c.head.node.number {
			target := &QueueChainNode{node, c.head}
			c.head = target
			cur = target
			continue
		}
		for cur.next != nil && cur.next.node.number < node.number {
			cur = cur.next
			count++
			if count > maxlength {
				break
			}
		}
		target := &QueueChainNode{node, cur.next}
		cur.next = target
		if count > maxlength {
			break
		}
	}
}

func (c *QueueChain) pop() *TreeNode1 {
	result := c.head
	c.head = result.next
	return result.node
}

type QueueChainNode struct {
	node *TreeNode1
	next *QueueChainNode
}
*/

func nthSuperUglyNumber(n int, primes []int) int {
	/*
		queue := QueueChain{&QueueChainNode{&TreeNode1{1, 0}, nil}}
		var res int
		for i := 0; i < n; i++ {
			node := queue.pop()
			res = node.number
			queue.insert(node.extend(primes), n-i)
			//fmt.Println(i)
		}
		return res
	*/
	res := make([]int, n)
	res[0] = 1
	pointers := make([]int, len(primes))
	for i := 1; i < n; i++ {
		min := -1
		minindex := 0
		for index, prime := range primes {
			res := prime * res[pointers[index]]
			switch {
			case min == -1:
				fallthrough
			case res < min:
				min = res
				minindex = index
			case res == min:
				pointers[index]++
			}
		}
		res[i] = min
		pointers[minindex]++
	}
	return res[n-1]
}
