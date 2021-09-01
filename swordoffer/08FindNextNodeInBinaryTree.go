package main

import "fmt"

/*
08. 二叉树的下一个节点
给定一颗二叉树和其中一个节点，如何找出中序遍历序列的下一个节点？
树中的节点除了有两个分别指向左、右子节点的指针，还有一个指向父节点的指针。
 */
/*
分析：
1、如果该节点有右子节点，则下一个节点为右子树的最左子节点
2、如果该节点是左子节点，则下一个节点为父节点
3、如果该节点不是左子节点，也没有右节点，则下一个节点为，第一个是左子节点的父节点
 */
func main() {
	a := TreeParentNode{Val: "a"}
	b := TreeParentNode{Val: "b"}
	c := TreeParentNode{Val: "c"}
	d := TreeParentNode{Val: "d"}
	e := TreeParentNode{Val: "e"}
	f := TreeParentNode{Val: "f"}
	g := TreeParentNode{Val: "g"}
	h := TreeParentNode{Val: "h"}
	i := TreeParentNode{Val: "i"}
	a.Left = &b
	a.Right = &c

	b.Left = &d
	b.Right = &e
	b.Parent = &a

	c.Left = &f
	c.Right = &g
	c.Parent = &a

	d.Parent = &b
	e.Parent = &b

	f.Parent = &c
	g.Parent = &c

	e.Left = &h
	e.Right = &i
	e.Parent = &b

	h.Parent = &e
	i.Parent = &e

	node := findNextNode(&b)
	if node != nil {
		fmt.Println(node.Val)
	}
}

func findNextNode(node *TreeParentNode) *TreeParentNode {
	if node == nil {
		return nil
	}
	// 1、如果该节点有右子节点，则下一个节点为右子树的最左子节点
	if node.Right != nil {
		return findLeftEndNode(node.Right)
	}
	if node.Parent == nil {
		return findLeftEndNode(node.Right)
	}
	// 2、如果该节点是左子节点，则下一个节点为父节点
	if node.Parent.Left == node {
		return node.Parent
	}
	// 3、如果该节点不是左子节点，也没有右节点，则下一个节点为，第一个是左子节点的父节点
	return findFirstLeftNodeParent(node)
}

func findFirstLeftNodeParent(node *TreeParentNode)  *TreeParentNode {
	for  {
		if node.Parent.Parent == nil {
			return nil
		}
		if node.Parent.Parent.Left == node.Parent {
			return node.Parent.Parent
		}
		node = node.Parent
	}

}

func findLeftEndNode(node *TreeParentNode)  *TreeParentNode{
	if node == nil {
		return nil
	}
	for node.Left != nil {
		node = node.Left
	}
	return node
}

type TreeParentNode struct {
	Val   string
	Left  *TreeParentNode
	Right *TreeParentNode
	Parent *TreeParentNode
}


