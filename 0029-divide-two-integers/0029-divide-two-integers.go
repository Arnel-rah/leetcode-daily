func divide(dividend int, divisor int) int {
    if dividend == -2147483648 && divisor == -1 {
        return 2147483647
    }

    neg := (dividend < 0) != (divisor < 0)
    a, b := abs(dividend), abs(divisor)

    res := 0
    for a >= b {
        k, bb := 1, b
        for a >= (bb << 1) {
            bb <<= 1
            k <<= 1
        }
        a -= bb
        res += k
    }

    if neg {
        res = -res
    }

    if res > 2147483647 {
        return 2147483647
    }
    if res < -2147483648 {
        return -2147483648
    }
    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}