package main

/**
10- II. 青蛙跳台阶问题
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
 */
func main() {

}

func FrogLoop(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var fib1 = 1
	var fib2 = 2
	var fibX int
	for i:=3; i<=n; i++ {
		fibX = (fib1 + fib2) % 1000000007
		fib1 = fib2
		fib2 = fibX
	}
	return fibX
}
