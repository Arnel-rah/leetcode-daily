const MOD = 1000000007

func concatenatedBinary(n int) int {
    result := 0
    length := 0

    for i := 1; i <= n; i++ {
        bits := bitLength(i)
        result = (result * pow2(bits, MOD) % MOD + i) % MOD
        length += bits
    }

    return result
}

func bitLength(x int) int {
    count := 0
    for x > 0 {
        count++
        x >>= 1
    }
    return count
}

func pow2(exp int, mod int) int {
    base := 2
    result := 1

    for exp > 0 {
        if exp&1 == 1 {
            result = (result * base) % mod
        }
        base = (base * base) % mod
        exp >>= 1
    }

    return result
}