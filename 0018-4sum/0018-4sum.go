func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    ans := [][]int{}

    for i := 0; i < n-3; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j := i + 1; j < n-2; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }

            l, r := j+1, n-1
            for l < r {
                s := nums[i] + nums[j] + nums[l] + nums[r]
                if s == target {
                    ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
                    for l < r && nums[l] == nums[l+1] {
                        l++
                    }
                    for l < r && nums[r] == nums[r-1] {
                        r--
                    }
                    l++
                    r--
                } else if s < target {
                    l++
                } else {
                    r--
                }
            }
        }
    }
    return ans
}