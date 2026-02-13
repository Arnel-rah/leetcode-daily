func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    for i := range haystack {
        if i+len(needle) > len(haystack) {
            break
        }
        if haystack[i:i+len(needle)] == needle {
            return i
        }
    }
    return -1
}