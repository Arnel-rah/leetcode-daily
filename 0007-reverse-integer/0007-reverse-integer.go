func reverse(x int) int {
    reversed := 0
    
    for x != 0 {
        pop := x % 10
        x /= 10
        if reversed > 214748364 || reversed < -214748364 {
            return 0
        }
        if reversed == 214748364 && pop > 7 {
            return 0
        }
        if reversed == -214748364 && pop < -8 {
            return 0
        }
        
        reversed = reversed*10 + pop
    }
    
    return reversed
}