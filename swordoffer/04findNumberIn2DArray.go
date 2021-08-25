package main

import "fmt"

/*
题目：04. 二维数组中的查找
在一个 n * m 的二维数组中，
每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数
 */
func main() {
	var matrix = [][]int{
		{1,2,8,9},
		{2,4,9,12},
		{4,7,10,13},
		{6,8,11,15},
	}
	fmt.Println(findNumberIn2DArray(matrix, 11))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	var j = len(matrix[0]) - 1
	for i:=0; i< len(matrix); i++ {
		for ; j>=0; j-- {
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] < target {
				break
			}
		}
	}
	return false
}
