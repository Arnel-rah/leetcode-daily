func longestConsecutive(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    
    sort.Ints(nums)
    
    maxLen := 1
    current := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            continue
        }
        if nums[i] == nums[i-1]+1 {
            current++
            maxLen = max(maxLen, current)
        } else {
            current = 1
        }
    }
    return maxLen
}