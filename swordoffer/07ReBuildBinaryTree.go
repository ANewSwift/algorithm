package main

import "fmt"

/*
07. 重建二叉树
输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
*/
func main() {
	var preorder = []int{1,2,4,7,3,5,6,8}
	var inorder = []int{4,7,2,1,5,3,8,6}
	head := buildTree(preorder, inorder)
	middleOrderPrintBinaryTree(head)
}

func middleOrderPrintBinaryTree(head *TreeNode) {
	if head == nil {
		return
	}
	middleOrderPrintBinaryTree(head.Left)
	fmt.Println(head.Val)
	middleOrderPrintBinaryTree(head.Right)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	index := 0
	for ; index < len(inorder); index++ {
		if preorder[0] == inorder[index] {
			break
		}
	}
	return &TreeNode{
		preorder[0],
		buildTree(preorder[getIndex(1,len(preorder)):getIndex(index+1,len(preorder))], inorder[0:index]),
		buildTree(preorder[getIndex(index+1,len(preorder)):], inorder[getIndex(index+1,len(inorder)):]),
	}
}

func getIndex(index int, len int) int {
	if index < 0 {
		return 0
	}
	if index > len {
		return len
	}
	return index
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
