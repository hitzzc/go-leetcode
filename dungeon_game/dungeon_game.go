package dungeon_game

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}
	matrix := make([][]int, len(dungeon))
	for i := range matrix {
		matrix[i] = make([]int, len(dungeon[0]))
	}
	rows := len(dungeon) - 1
	cols := len(dungeon[0]) - 1
	matrix[rows][cols] = dungeon[rows][cols]
	for i := rows - 1; i >= 0; i-- {
		matrix[i][cols] = min(dungeon[i][cols], matrix[i+1][cols]+dungeon[i][cols])
	}
	for j := cols - 1; j >= 0; j-- {
		matrix[rows][j] = min(dungeon[rows][j], matrix[rows][j+1]+dungeon[rows][j])
	}
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			matrix[i][j] = min(dungeon[i][j], dungeon[i][j]+max(matrix[i+1][j], matrix[i][j+1]))
		}
	}
	if matrix[0][0] >= 0 {
		return 1
	}
	return 1 - 1*matrix[0][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
