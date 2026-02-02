func myAtoi(s string) int {
    i, sign, result := 0, 1, 0
    n := len(s)

    for i < n && s[i] == ' ' {
        i++
    }
    if i == n {
        return 0
    }
    if s[i] == '-' {
        sign = -1
        i++
    } else if s[i] == '+' {
        i++
    }

    for i < n && s[i] >= '0' && s[i] <= '9' {
        digit := int(s[i] - '0')

        if result > (2147483647-digit)/10 {
            if sign == 1 {
                return 2147483647
            }
            return -2147483648
        }

        result = result*10 + digit
        i++
    }

    return sign * result
}