func letterCombinations(digits string) []string {
    if digits == "" {
        return nil
    }

    m := map[byte]string{
        '2': "abc", '3': "def", '4': "ghi", '5': "jkl",
        '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
    }

    var ans []string

    var dfs func(int, string)
    dfs = func(i int, s string) {
        if i == len(digits) {
            ans = append(ans, s)
            return
        }
        for _, c := range m[digits[i]] {
            dfs(i+1, s+string(c))
        }
    }

    dfs(0, "")
    return ans
}