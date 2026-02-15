func firstMissingPositive(nums []int) int {
    n := len(nums)
    
    for i := 0; i < n; i++ {
        if nums[i] <= 0 || nums[i] > n {
            nums[i] = n + 1
        }
    }
    
    for i := 0; i < n; i++ {
        val := abs(nums[i])
        if val <= n {
            nums[val-1] = -abs(nums[val-1])
        }
    }
    
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            return i + 1
        }
    }
    
    return n + 1
}

func abs(x int) int {
    if x < 0 { return -x }
    return x
}