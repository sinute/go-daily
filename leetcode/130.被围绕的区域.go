package leetcode

/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
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
		dfs130(board, i, 0)
		dfs130(board, i, cols-1)
	}
	for j := 0; j < cols; j++ {
		dfs130(board, 0, j)
		dfs130(board, rows-1, j)
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

func dfs130(board [][]byte, x, y int) {
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
	dfs130(board, x-1, y)
	dfs130(board, x+1, y)
	dfs130(board, x, y-1)
	dfs130(board, x, y+1)
}

// @lc code=end
