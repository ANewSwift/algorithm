package main

import "fmt"

/*
	题目03.找出数组中重复的数字
	在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
	数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
	请找出数组中任意一个重复的数字。
	限制：2 <= n <= 100000
 */
func main()  {
	number := findRepeatNumber([]int{3, 4, 2, 0, 0, 1})
	fmt.Println(number)
}

/*
解题思路：
	必须遍历到数值数组的每个值，然后把他们放到值对应的下标上，
	如果下标i == 下标对应的数值，则遍历下一个值
	如果不相等，则比较该下标值 与 对应的以改值为下标的数值
		如果相等，则返回该下标值
		如果不相等，则交换，（此时改下标值已放到正确位置），
			在保证该下标i与交换回来了的值相等时，遍历下一个值
			不相等则，拿该下标i，接着比较。
 */

func findRepeatNumber(nums []int) int {
	for index,num := range nums {
		for index != nums[index] {
			// target是该index位置的原始值
			var target = nums[index]
			if nums[index] == nums[target]{
				return target
			}
			var tmp = nums[target]
			nums[target] = num
			nums[index] = tmp
		}
	}
	return -1
}

