package surrounded_regions

type solution struct {
	visited [][]bool
	board   [][]byte
}

func solve(board [][]byte) {
	if len(board) < 2 {
		return
	}
	s := &solution{
		visited: make([][]bool, len(board)),
		board:   board,
	}
	for i := range board {
		s.visited[i] = make([]bool, len(board[i]))
	}
	for i := range board {
		if board[i][0] == 'O' {
			s.DFS(i, 0)
		}
		if len(board[i]) > 1 && board[i][len(board[i])-1] == 'O' {
			s.DFS(i, len(board[i])-1)
		}
	}
	for j := range board[0] {
		if board[0][j] == 'O' {
			s.DFS(0, j)
		}
	}
	for j := range board[len(board)-1] {
		if board[len(board)-1][j] == 'O' {
			s.DFS(len(board)-1, j)
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' && !s.visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
	return
}

func (this *solution) DFS(i, j int) {
	if i < 0 || i >= len(this.board) || j < 0 || j >= len(this.board[i]) || this.board[i][j] == 'X' || this.visited[i][j] {
		return
	}
	this.visited[i][j] = true
	this.DFS(i-1, j)
	this.DFS(i+1, j)
	this.DFS(i, j-1)
	this.DFS(i, j+1)
	return
}
