package main

import (
	"fmt"
	"strings"
)

/*
05. 替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."
 */
func main() {
	space := replaceSpace("We are happy.")
	fmt.Println(space)
}
func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
