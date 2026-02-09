func maxSubArray(nums []int) int {
    best, sum := nums[0], nums[0]
    for _, n := range nums[1:] {
        sum = max(n, sum+n)
        best = max(best, sum)
    }
    return best
}

func max(a, b int) int { if a > b { return a }; return b }