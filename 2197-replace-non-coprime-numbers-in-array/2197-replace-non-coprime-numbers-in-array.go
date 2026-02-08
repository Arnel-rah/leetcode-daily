func replaceNonCoprimes(nums []int) []int {
    stack := make([]int, 0, len(nums))
    
    for _, x := range nums {
        for len(stack) > 0 {
            y := stack[len(stack)-1]
            g := gcd(y, x)
            if g == 1 {
                break
            }
            x = int(int64(y) / int64(g) * int64(x))
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, x)
    }
    
    return stack
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}