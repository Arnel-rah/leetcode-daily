func generateParenthesis(n int) (res []string) {
    var f func(string, int, int)
    f = func(s string, l, r int) {
        if l == n && r == n {
            res = append(res, s)
            return
        }
        if l < n {
            f(s+"(", l+1, r)
        }
        if r < l {
            f(s+")", l, r+1)
        }
    }
    f("", 0, 0)
    return
}