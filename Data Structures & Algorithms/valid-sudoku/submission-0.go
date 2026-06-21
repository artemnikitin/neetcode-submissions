func isValidSudoku(board [][]byte) bool {
rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = map[byte]bool{}
		columns[i] = map[byte]bool{}
		squares[i] = map[byte]bool{}
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			value := board[r][c]
			squareIndex := (r/3)*3 + c/3
			if rows[r][value] || columns[c][value] || squares[squareIndex][value] {
				return false
			}
			rows[r][value] = true
			columns[c][value] = true
			squares[squareIndex][value] = true
		}
	}

	return true
}
