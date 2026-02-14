func isValid(s string) bool {
    stack := []rune{}
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, c := range s {
        if c == '(' || c == '[' || c == '{' {
            stack = append(stack, c)
        } else if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
            return false
        } else {
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}