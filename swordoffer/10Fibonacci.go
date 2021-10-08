package main

/**
10- I. 斐波那契数列
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下
F(0) = 0,
F(1) = 1,
F(n) = F(n-1) + F(n-2)
 */
func main() {

}

// 循环方式求解
func FibLoop(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var fib1 = 0
	var fib2 = 1
	var fibN int
	for i:=2; i<=n; i++ {
		fibN = (fib1 + fib2) % 1000000007
		fib1 = fib2
		fib2 = fibN
	}
	return fibN
}


// 递归方式求解
func FibRecursionHandle(n int) int  {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FibRecursion(0,1,2, n)
}

func FibRecursion(fib1, fib2, fibCur, fibN int) int {
	var fibX = (fib1 + fib2) % 1000000007
	fib1 = fib2
	fib2 = fibX
	if fibCur == fibN {
		return fibX
	}
	return FibRecursion(fib1, fib2, fibCur+1, fibN)
}