package valid_sudoku

import (
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	println(isValidSudoku(
		[][]byte{[]byte(".87654321"),
			[]byte("2........"),
			[]byte("3........"),
			[]byte("4........"),
			[]byte("5........"),
			[]byte("6........"),
			[]byte("7........"),
			[]byte("8........"),
			[]byte("9........")}))
}
