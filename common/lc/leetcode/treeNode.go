package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTheTree(treeNode *TreeNode) {
	fmt.Printf("[")
	printNode(treeNode)
	fmt.Printf("]")
}
func printNode(treeNode *TreeNode) {
	fmt.Printf("%d,", treeNode.Val)
	if treeNode.Left == nil {
		fmt.Printf("%s,", "null")
	} else {
		printNode(treeNode.Left)
	}
	if treeNode.Right == nil {
		fmt.Printf("%s,", "null")
	} else {
		printNode(treeNode.Right)
	}
}
