package main

import "fmt"

//130. 被围绕的区域
func solve(board [][]byte) {
    rows := len(board)
    if rows <= 2 {
        return
    }
    cols := len(board[0])
    if cols <= 2 {
        return
    }
    for i := 0; i < rows; i++ {
        dfs(board, i, 0)
        dfs(board, i, cols-1)
    }
    for j := 0; j < cols; j++ {
        dfs(board, 0, j)
        dfs(board, rows-1, j)
    }
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            } else if board[i][j] == '*' {
                board[i][j] = 'O'
            }
        }
    }
}

func dfs(board [][]byte, x, y int) {
    rows := len(board)
    cols := len(board[0])
    if x == rows || x < 0 {
        return
    }
    if y == cols || y < 0 {
        return
    }
    if board[x][y] != 'O' {
        return
    }
    board[x][y] = '*'
    dfs(board, x-1, y)
    dfs(board, x+1, y)
    dfs(board, x, y-1)
    dfs(board, x, y+1)
}

func main() {
    t := [][]byte{
        {'X', 'X', 'X', 'X'},
        {'X', 'O', 'O', 'X'},
        {'X', 'X', 'O', 'X'},
        {'X', 'O', 'X', 'X'},
    }
    solve(t)
    fmt.Println(t)

    t = [][]byte{
        {'O', 'O', 'O', 'O'},
        {'O', 'O', 'O', 'O'},
        {'O', 'O', 'O', 'O'},
        {'O', 'O', 'O', 'O'},
    }
    solve(t)
    fmt.Println(t)

    t = [][]byte{
        {'X', 'O', 'X', 'O', 'X', 'O'},
        {'O', 'X', 'O', 'X', 'O', 'X'},
        {'X', 'O', 'O', 'X', 'X', 'O'},
        {'O', 'X', 'O', 'X', 'O', 'X'},
    }
    solve(t)
    fmt.Println(t)

    t = [][]byte{
        {'O', 'X', 'X', 'O', 'X'},
        {'X', 'O', 'O', 'X', 'O'},
        {'X', 'O', 'X', 'O', 'X'},
        {'O', 'X', 'O', 'O', 'O'},
        {'X', 'X', 'O', 'X', 'O'},
    }
    solve(t)
    fmt.Println(t)
}

//输入：
//[["O","X","X","O","X"],["X","O","O","X","O"],["X","O","X","O","X"],["O","X","O","O","O"],["X","X","O","X","O"]]
//输出：
//[["O","X","X","O","X"],["X","X","X","X","O"],["X","X","X","X","X"],["O","X","X","X","O"],["X","X","O","X","O"]]
//预期结果：
//[["O","X","X","O","X"],["X","X","X","X","O"],["X","X","X","O","X"],["O","X","O","O","O"],["X","X","O","X","O"]]