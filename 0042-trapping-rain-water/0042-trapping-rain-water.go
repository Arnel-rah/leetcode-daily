func trap(height []int) int {
    l, r := 0, len(height)-1
    lMax, rMax, ans := 0, 0, 0

    for l < r {
        if height[l] <= height[r] {
            if height[l] >= lMax {
                lMax = height[l]
            } else {
                ans += lMax - height[l]
            }
            l++
        } else {
            if height[r] >= rMax {
                rMax = height[r]
            } else {
                ans += rMax - height[r]
            }
            r--
        }
    }
    return ans
}