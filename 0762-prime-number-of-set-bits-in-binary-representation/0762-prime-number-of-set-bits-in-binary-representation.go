func countPrimeSetBits(left int, right int) int {
    isPrime := [21]bool{2: true, 3: true, 5: true, 7: true,
                        11: true, 13: true, 17: true, 19: true}

    count := 0
    for i := left; i <= right; i++ {
        bits := 0
        n := i
        for n > 0 {
            bits += n & 1
            n >>= 1
        }
        if isPrime[bits] {
            count++
        }
    }
    return count
}