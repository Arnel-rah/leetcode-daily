func jump(nums []int) int {
    jumps, farthest, end := 0, 0, 0
    for i := 0; i < len(nums)-1; i++ {
        farthest = max(farthest, i+nums[i])
        if i == end {
            jumps++
            end = farthest
        }
    }
    return jumps
}

func max(a, b int) int { if a > b { return a }; return b }