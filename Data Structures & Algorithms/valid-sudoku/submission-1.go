func isValidSudoku(board [][]byte) bool {
rows := make([]int, 9)
	columns := make([]int, 9)
	squares := make([]int, 9)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			value := board[r][c] - '1'
			bit := 1 << value
			squareIndex := (r/3)*3 + c/3

			if rows[r]&bit != 0 || columns[c]&bit != 0 ||
				squares[squareIndex]&bit != 0 {
				return false
			}

			rows[r] |= bit
			columns[c] |= bit
			squares[squareIndex] |= bit
		}
	}

	return true
}
