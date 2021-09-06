package main

import "fmt"

/*
09. 用两个栈实现队列
用两个栈实现一个队列。
队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。
(若队列中没有元素，deleteHead操作返回 -1 )
 */
func main() {

}


func (this *CQueue) AppendTail(value int) {
	this.Q2.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.Q1.Tail != nil {
		return this.Q1.Pop()
	}
	for this.Q2.Tail != nil {
		this.Q1.Push(this.Q2.Pop())
	}
	if this.Q1.Tail == nil {
		return -1
	}
	return this.Q1.Pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
func Constructor() CQueue {
	return CQueue{
		&Queue{nil},
		&Queue{nil},
	}
}


type CQueue struct {
	Q1 *Queue
	Q2 *Queue
}

type Queue struct {
	Tail *Node
}

func (queue *Queue) Push(value int) {
	queue.Tail = &Node{value, queue.Tail}
}

func (queue *Queue) Pop() int {
	if queue.Tail == nil {
		return -1
	}
	node := queue.Tail
	queue.Tail = queue.Tail.Next
	node.Next = nil
	return node.Val
}

func (queue *Queue) Print() {
	point := queue.Tail
	for point != nil {
		fmt.Println(point.Val)
		point = point.Next
	}
}

type Node struct {
	Val int
	Next *Node
}
