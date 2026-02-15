func addBinary(a, b string) string {
    res := ""
    carry := 0
    i, j := len(a)-1, len(b)-1
    
    for i >= 0 || j >= 0 || carry > 0 {
        sum := carry
        if i >= 0 {
            sum += int(a[i] - '0')
            i--
        }
        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }
        res = string('0'+sum%2) + res
        carry = sum / 2
    }
    
    return res
}