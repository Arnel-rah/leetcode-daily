func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    used := make([]bool, len(nums))
    
    var backtrack func(path []int)
    backtrack = func(path []int) {
        if len(path) == len(nums) {
            res = append(res, append([]int{}, path...))
            return
        }
        
        for i := 0; i < len(nums); i++ {
            // sauter les doublons
            if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
                continue
            }
            
            used[i] = true
            backtrack(append(path, nums[i]))
            used[i] = false
        }
    }
    
    backtrack([]int{})
    return res
}