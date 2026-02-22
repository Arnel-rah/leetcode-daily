func climbStairs(n int) int {
    memo := make([]int, n+1)
    return helper(n, memo)
}

func helper(n int, memo []int) int {
    if n <= 2 {
        return n
    }
    if memo[n] != 0 {
        return memo[n]
    }
    memo[n] = helper(n-1, memo) + helper(n-2, memo)
    return memo[n]
}