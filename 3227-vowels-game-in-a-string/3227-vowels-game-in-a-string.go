func doesAliceWin(s string) bool {
    for i := range s {
        switch s[i] {
        case 'a','e','i','o','u':
            return true
        }
    }
    return false
}
