package fibonacci_sequence

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

func recursionClimbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	return recursionClimbStairs(n-1) + recursionClimbStairs(n-2)
}

func removeDupRecursionClimbStairs(n int) int {
	memo := make([]int, n+1)
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 { // 递归边界
			return 1
		}
		p := &memo[i]
		if *p != 0 { // 之前计算过
			return *p
		}
		res := dfs(i-1) + dfs(i-2)
		*p = res // 记忆化
		return res
	}
	return dfs(n)
}

func iterativeClimbStairs(n int) int {
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
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
