func permute(nums []int) [][]int {
    res := [][]int{}
    
    var dfs func(int)
    dfs = func(pos int) {
        if pos == len(nums) {
            res = append(res, append([]int{}, nums...))
            return
        }
        
        for i := pos; i < len(nums); i++ {
            nums[pos], nums[i] = nums[i], nums[pos]
            dfs(pos + 1)
            nums[pos], nums[i] = nums[i], nums[pos]
        }
    }
    
    dfs(0)
    return res
}