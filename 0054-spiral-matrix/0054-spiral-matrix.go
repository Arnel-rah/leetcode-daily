func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return nil
    }

    res := []int{}
    m, n := len(matrix), len(matrix[0])
    r1, r2, c1, c2 := 0, m-1, 0, n-1

    for r1 <= r2 && c1 <= c2 {
        for c := c1; c <= c2; c++ { res = append(res, matrix[r1][c]) }
        r1++
        for r := r1; r <= r2; r++ { res = append(res, matrix[r][c2]) }
        c2--
        if r1 <= r2 {
            for c := c2; c >= c1; c-- { res = append(res, matrix[r2][c]) }
            r2--
        }
        if c1 <= c2 {
            for r := r2; r >= r1; r-- { res = append(res, matrix[r][c1]) }
            c1++
        }
    }

    return res
}