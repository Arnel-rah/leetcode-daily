

func minScoreTriangulation(values []int) int {
    n := len(values)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    for length := 2; length < n; length++ {
        for i := 0; i < n-length; i++ {
            j := i + length
            dp[i][j] = math.MaxInt32

            for k := i + 1; k < j; k++ {
                cost := values[i]*values[k]*values[j] + dp[i][k] + dp[k][j]
                if cost < dp[i][j] {
                    dp[i][j] = cost
                }
            }
        }
    }

    return dp[0][n-1]
}