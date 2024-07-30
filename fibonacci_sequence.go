package dp

//The Fibonacci sequence pattern is useful when the solution to
//a problem depends on the solutions of smaller instances of the same problem.

//There is a clear recursive relationship,
//often resembling the classic Fibonacci sequence F(n) = F(n-1) + F(n-2).

//[LeetCode 746. Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/description/)

// [LeetCode 70: Climbing Stairs](https://leetcode.com/problems/climbing-stairs/description/)
func climbStairs(n int) int {
	var a, b = 1, 1
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
	}
	return b
}

// [LeetCode 509: Fibonacci Number](https://leetcode.com/problems/fibonacci-number/description/)
func fib(n int) int {
	//f(n) = f(n-1)+f(n-2)
	var n1, n2 = 0, 1
	for i := 0; i < n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n1
}
