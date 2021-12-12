package main

import (
	"fmt"
)
/*
快速排序：
挑选最后一个标志，小的放左边，大的放右边
分别对左右两部分，做递归
 */
func main() {
	ints := []int{3, 1, 5, 4, 2}
	quickSort(ints)
	fmt.Println(ints)
}

func quickSort(slice []int) {
	if len(slice) == 0 || len(slice) == 1 {
		return
	}
	mid := len(slice) - 1
	for i:=0; i<len(slice); i++ {
		if slice[i] > slice[mid] && i < mid {
			slice[i],slice[mid] = slice[mid], slice[i]
			mid = i
		}
		if slice[i] < slice[mid] && i > mid {
			slice[i],slice[mid] = slice[mid], slice[i]
			mid = i
		}

	}
	quickSort(slice[:mid])
	if mid < len(slice) {
		quickSort(slice[mid:])
	}
}
