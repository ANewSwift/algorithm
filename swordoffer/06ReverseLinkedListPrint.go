package main

import "fmt"

/*
06. 从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
 */

func main() {
	head := &ListNode{0,nil}
	var p = head
	for i:=1; i<10; i++ {
		p.Next = &ListNode{i,nil}
		p = p.Next
	}
	ints := reversePrint(head)
	fmt.Println(ints)
}

func reversePrint(head *ListNode) []int {
	var p = head
	count := 0
	for p != nil {
		count++
		p = p.Next
	}
	p = head
	rev := make([]int, count)
	for p  != nil {
		count--
		rev[count] = p.Val
		p = p.Next
	}
	return rev
}

type ListNode struct {
	Val  int
	Next *ListNode
}


