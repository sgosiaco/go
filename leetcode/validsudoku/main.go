func isValidSudoku(board [][]byte) bool {
	colMaps := [9][9]bool{}
	boxMaps := [3][3][9]bool{}

	for y, row := range board {
		rowMap := [9]bool{}
		for x, col := range row {
			if string(col) != "." {
				index := int(col) - 49
				if rowMap[index] {
					return false
				}
				if colMaps[x][index] {
					return false
				}
				if boxMaps[x/3][y/3][index] {
					return false
				}
				rowMap[index] = true
				colMaps[x][index] = true
				boxMaps[x/3][y/3][index] = true
			}

		}
	}
	return true
}