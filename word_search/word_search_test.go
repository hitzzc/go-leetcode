package word_search

import (
	"testing"
)

func TestExist(t *testing.T) {
	board := [][]byte{
		[]byte("abce"),
		[]byte("sfcs"),
		[]byte("adee"),
	}
	println(exist(board, "abceseeefs"))
}
