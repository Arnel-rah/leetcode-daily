func isValidSudoku(board [][]byte) bool {
    var row, col, box [9][9]bool

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == '.' {
                continue
            }
            num := board[i][j] - '1'
            k := (i/3)*3 + j/3

            if row[i][num] || col[j][num] || box[k][num] {
                return false
            }

            row[i][num], col[j][num], box[k][num] = true, true, true
        }
    }
    return true
}