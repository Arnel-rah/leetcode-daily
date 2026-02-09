func solveNQueens(n int) [][]string {
    var res [][]string
    board := make([][]byte, n)
    for i := range board {
        board[i] = make([]byte, n)
        for j := range board[i] {
            board[i][j] = '.'
        }
    }

    col, diag1, diag2 := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1)

    var dfs func(int)
    dfs = func(r int) {
        if r == n {
            sol := make([]string, n)
            for i := range board {
                sol[i] = string(board[i])
            }
            res = append(res, sol)
            return
        }
        for c := 0; c < n; c++ {
            if col[c] || diag1[r+c] || diag2[r-c+n-1] {
                continue
            }
            board[r][c] = 'Q'
            col[c], diag1[r+c], diag2[r-c+n-1] = true, true, true
            dfs(r + 1)
            board[r][c] = '.'
            col[c], diag1[r+c], diag2[r-c+n-1] = false, false, false
        }
    }

    dfs(0)
    return res
}