func searchMatrix(matrix [][]int, target int) bool {
	targetRow := len(matrix) - 1
	for i := 1; i < len(matrix); i++ {
		if target < matrix[i][0] {
			targetRow = i - 1
			break
		}
	}
	for _, v := range matrix[targetRow] {
		if target == v {
			return true
		}
	}

	return false
}