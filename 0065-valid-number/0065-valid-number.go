func isNumber(s string) bool {
    seenDigit := false
    seenDot := false
    seenExponant := false

    for i := 0; i < len(s); i++ {
        char := s[i]

        switch {
        case char >= '0' && char <= '9':
            seenDigit = true

        case char == '+' || char == '-':
            if i > 0 && s[i-1] != 'e' && s[i-1] != 'E' {
                return false
            }

        case char == 'e' || char == 'E':
            if seenExponant || !seenDigit {
                return false
            }
            seenExponant = true
            seenDigit = false

        case char == '.':
            if seenDot || seenExponant {
                return false
            }
            seenDot = true

        default:
            return false
        }
    }

    return seenDigit
}