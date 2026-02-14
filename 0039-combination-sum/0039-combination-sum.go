func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    var f func(int, int, []int)
    f = func(i, t int, p []int) {
        if t == 0 {
            res = append(res, append([]int{}, p...))
            return
        }
        if t < 0 || i == len(candidates) {
            return
        }
        f(i, t-candidates[i], append(p, candidates[i]))
        f(i+1, t, p)
    }
    f(0, target, []int{})
    return res
}