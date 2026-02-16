func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    
    for _, s := range strs {
        key := sortString(s)
        m[key] = append(m[key], s)
    }
    
    ans := make([][]string, 0, len(m))
    for _, v := range m {
        ans = append(ans, v)
    }
    return ans
}

func sortString(s string) string {
    b := []byte(s)
    sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
    return string(b)
}