func hasIncreasingSubarrays(nums []int, k int) bool {
    n := len(nums)

    for a := 0; a+2*k <= n; a++ {

        inc1 := true
        for i := a; i < a+k-1; i++ {
            if nums[i] >= nums[i+1] {
                inc1 = false
                break
            }
        }

        inc2 := true
        for i := a+k; i < a+2*k-1; i++ {
            if nums[i] >= nums[i+1] {
                inc2 = false
                break
            }
        }

        if inc1 && inc2 {
            return true
        }
    }

    return false
}