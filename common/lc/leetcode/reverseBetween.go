package leetcode

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	return reverseBetween(head, left, right)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	ls := head
	ll := left - 1
	idx := 1
	var newNode, leftNode, firstLeftNode *ListNode
	for ls != nil {
		if idx == ll {
			leftNode = ls
		}
		if idx == left {
			firstLeftNode = ls
			newNode = ls
		}

		if idx == right+1 {
			break
		}
		//已经出现了left 开始头部插入
		if newNode != nil {
			tmpNode := ls
			ls = ls.Next
			tmpNode.Next = newNode
			newNode = tmpNode
			idx++
			continue
		}
		ls = ls.Next
		idx++
	}
	if leftNode != nil {
		leftNode.Next = newNode
	} else { //这个时候没有头啊 是从第一个开始的
		head = newNode
	}
	if firstLeftNode != nil {
		firstLeftNode.Next = ls
	}
	return head
}
