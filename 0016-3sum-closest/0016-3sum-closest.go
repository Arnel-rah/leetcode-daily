func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    n := len(nums)
    closest := nums[0] + nums[1] + nums[2]

    for i := 0; i < n-2; i++ {
        l, r := i+1, n-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if abs(sum-target) < abs(closest-target) {
                closest = sum
            }
            if sum < target {
                l++
            } else {
                r--
            }
        }
    }
    return closest
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}