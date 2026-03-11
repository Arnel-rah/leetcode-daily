func maxIncreasingSubarrays(nums []int) int {
    n := len(nums)
    runs := []int{}
    
    length := 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            length++
        } else {
            runs = append(runs, length)
            length = 1
        }
    }
    runs = append(runs, length)

    ans := 0

    for i := 0; i < len(runs); i++ {
        if runs[i]/2 > ans {
            ans = runs[i] / 2
        }

        if i > 0 {
            m := runs[i]
            if runs[i-1] < m {
                m = runs[i-1]
            }
            if m > ans {
                ans = m
            }
        }
    }

    return ans
}