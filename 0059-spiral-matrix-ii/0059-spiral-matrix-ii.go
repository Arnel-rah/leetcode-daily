func generateMatrix(n int) [][]int {
    // 1. Initialisation de la matrice n x n
    matrix := make([][]int, n)
    for i := range matrix {
        matrix[i] = make([]int, n)
    }

    left, right := 0, n-1
    top, bottom := 0, n-1
    num := 1
    target := n * n

    for num <= target {
        for i := left; i <= right; i++ {
            matrix[top][i] = num
            num++
        }
        top++

        for i := top; i <= bottom; i++ {
            matrix[i][right] = num
            num++
        }
        right--

        for i := right; i >= left; i-- {
            matrix[bottom][i] = num
            num++
        }
        bottom--

        for i := bottom; i >= top; i-- {
            matrix[i][left] = num
            num++
        }
        left++
    }

    return matrix
}