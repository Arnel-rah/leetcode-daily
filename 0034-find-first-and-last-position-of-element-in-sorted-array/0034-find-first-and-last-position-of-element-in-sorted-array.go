func searchRange(nums []int, target int) []int {
    first := binarySearch(nums, target, true)
    if first == -1 {
        return []int{-1, -1}
    }
    last := binarySearch(nums, target, false)
    return []int{first, last}
}

func binarySearch(nums []int, target int, findFirst bool) int {
    left, right := 0, len(nums)-1
    ans := -1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            ans = mid
            if findFirst {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return ans
}