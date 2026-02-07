func removeDuplicates(nums []int) int {
    k := 0
    for _, num := range nums {
        if k == 0 || num != nums[k-1] {
            nums[k] = num
            k++
        }
    }
    return k
}