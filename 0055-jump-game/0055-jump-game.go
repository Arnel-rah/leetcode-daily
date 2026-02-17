func canJump(nums []int) bool {
    maxReach := 0
    for i := 0; i <= maxReach && i < len(nums); i++ {
        if i+nums[i] > maxReach {
            maxReach = i + nums[i]
        }
    }
    return maxReach >= len(nums)-1
}