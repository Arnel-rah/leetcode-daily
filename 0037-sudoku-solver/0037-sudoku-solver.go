func solveSudoku(board [][]byte) {
    var row, col, box [9]int

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] != '.' {
                num := 1 << (board[i][j] - '1')
                k := (i/3)*3 + j/3
                row[i] |= num
                col[j] |= num
                box[k] |= num
            }
        }
    }

    var dfs func(int, int) bool
    dfs = func(r, c int) bool {
        if r == 9 {
            return true
        }
        nr, nc := r, c+1
        if nc == 9 {
            nr, nc = r+1, 0
        }

        if board[r][c] != '.' {
            return dfs(nr, nc)
        }

        k := (r/3)*3 + c/3
        for num := 0; num < 9; num++ {
            mask := 1 << num
            if (row[r]&mask)|(col[c]&mask)|(box[k]&mask) == 0 {
                board[r][c] = byte('1' + num)
                row[r] |= mask
                col[c] |= mask
                box[k] |= mask

                if dfs(nr, nc) {
                    return true
                }

                board[r][c] = '.'
                row[r] &= ^mask
                col[c] &= ^mask
                box[k] &= ^mask
            }
        }
        return false
    }

    dfs(0, 0)
}