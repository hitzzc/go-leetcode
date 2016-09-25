package search_a_2D_matrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := -1
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] <= target && (i < len(matrix)-1 && matrix[i+1][0] > target || i == len(matrix)-1) {
			row = i
			break
		}
	}
	if row == -1 {
		return false
	}
	return binarySearch(matrix[row], 0, len(matrix[row])-1, target)
}

func binarySearch(array []int, i, j, target int) bool {
	if i > j {
		return false
	}
	mid := i + (j-i)/2
	if array[mid] == target {
		return true
	}
	success := binarySearch(array, i, mid-1, target)
	if success {
		return true
	}
	return binarySearch(array, mid+1, j, target)
}
