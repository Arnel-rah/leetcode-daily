func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    
    var res [][]int
    
    var backtrack func(start int, remain int, path []int)
    backtrack = func(start int, remain int, path []int) {
        if remain == 0 {
            res = append(res, append([]int{}, path...))
            return
        }
        if remain < 0 {
            return
        }
        
        prev := -1
        for i := start; i < len(candidates); i++ {
            if candidates[i] == prev {
                continue // saute les doublons
            }
            if candidates[i] > remain {
                break
            }
            
            backtrack(i+1, remain-candidates[i], append(path, candidates[i]))
            prev = candidates[i]
        }
    }
    
    backtrack(0, target, []int{})
    return res
}