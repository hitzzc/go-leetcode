package word_search

func exist(board [][]byte, word string) bool {
	words := []byte(word)
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return true
	}

	mask := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		mask[i] = make([]int, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

		}
	}
	return false
}

func getExist(board [][]byte, words []byte, index int, row int, col int, mask [][]int) {

}
