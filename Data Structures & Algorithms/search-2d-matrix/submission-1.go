func searchMatrix(matrix [][]int, target int) bool {
    first := 0
	last := len(matrix) - 1
	row := -1
	for first <= last {
		mid := first + (last-first)/2
		if matrix[mid][0] <= target && target <= matrix[mid][len(matrix[mid])-1] {
			row = mid
			break
		} else if matrix[mid][0] > target {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	if row == -1 {
		return false
	}

	first = 0
	last = len(matrix[row]) - 1
	for first <= last {
		mid := first + (last-first)/2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] > target {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}

	return false
}
