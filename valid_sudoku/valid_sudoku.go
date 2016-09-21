package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	rowMask, colMask, areaMask := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for r := range board {
		for c := range board[r] {
			if board[r][c] == '.' {
				continue
			}
			digit := board[r][c] - '0' - 1
			area := 3*(r/3) + c/3
			if rowMask[r][digit] || colMask[c][digit] || areaMask[area][digit] {
				return false
			}
			rowMask[r][digit] = true
			colMask[c][digit] = true
			areaMask[area][digit] = true
		}
	}
	return true
}
