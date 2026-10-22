func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
        dp[i] = -1
    }
	return helper(n, dp)
}

func helper(n int, dp []int) int {
    if n == 1 || n == 2 {
        return n
    }
    if dp[n] != -1 {
        return dp[n]
    } 
    dp[n] = helper(n - 1, dp) + helper(n - 2, dp)
    return dp[n]
}