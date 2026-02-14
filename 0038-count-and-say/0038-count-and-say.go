func countAndSay(n int) string {
    s := "1"
    for i := 1; i < n; i++ {
        next := ""
        for j := 0; j < len(s); {
            count := 1
            for j+count < len(s) && s[j+count] == s[j] {
                count++
            }
            next += string(byte(count+'0')) + string(s[j])
            j += count
        }
        s = next
    }
    return s
}