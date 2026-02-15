func isMatch(s string, p string) bool {
    dp := make([]bool, len(p)+1)
    dp[0] = true
    
    for j := 1; j <= len(p); j++ {
        if p[j-1] == '*' {
            dp[j] = dp[j-1]
        }
    }
    
    for i := 1; i <= len(s); i++ {
        prev := dp[0]
        dp[0] = false
        
        for j := 1; j <= len(p); j++ {
            temp := dp[j]
            if p[j-1] == '*' {
                dp[j] = dp[j-1] || temp
            } else if p[j-1] == '?' || p[j-1] == s[i-1] {
                dp[j] = prev
            } else {
                dp[j] = false
            }
            prev = temp
        }
    }
    
    return dp[len(p)]
}