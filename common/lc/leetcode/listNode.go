package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToNode(arr []int) *ListNode {
	l := &ListNode{}
	ls := l
	for i := 0; i < len(arr); i++ {
		nl := &ListNode{
			Val: arr[i],
		}
		l.Next = nl
		l = l.Next
	}
	return ls.Next
}

func PrintNode(l *ListNode) {
	ls := l
	arr := make([]int, 0)
	for ls != nil {
		arr = append(arr, ls.Val)
		ls = ls.Next
	}
	fmt.Println(arr)
}
