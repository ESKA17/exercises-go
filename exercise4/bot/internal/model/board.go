package model

import "fmt"

type Token string

const (
	TokenEmpty Token = " "
	TokenX     Token = "x"
	TokenO     Token = "o"
)

const (
	Cols = 3
	Rows = 3
)

type Board [Cols * Rows]Token

func (b *Board) CalculateNewPosition(token Token) int {
	ind := b.calculateNewPosition(token)
	b.logMove(ind, token)
	return ind
}

func (b *Board) logMove(pos int, token Token) {
	b[pos] = token
	fmt.Println(b.String())
}

func (b *Board) String() string {
	str := fmt.Sprintf(
		`
	%s|%s|%s
	-----
	%s|%s|%s
	-----
	%s|%s|%s
`,
		b[0],
		b[1],
		b[2],
		b[3],
		b[4],
		b[5],
		b[6],
		b[7],
		b[8],
	)

	return str
}

func (b *Board) calculateNewPosition(token Token) int {
	newBoard := b
	var bestMove int
	isMax := token == TokenX
	first := isFirstMove(newBoard)

	if first {
		return 4
	}

	if isMax {
		best := -1000
		for i := 0; i < 9; i++ {
			if newBoard[i] == TokenEmpty {
				newBoard[i] = TokenX
				score := minimax(newBoard, 0, false)
				newBoard[i] = TokenEmpty

				if score > best {
					best = score
					bestMove = i
				}
			}
		}
	} else {
		best := 1000
		for i := 0; i < 9; i++ {
			if newBoard[i] == TokenEmpty {
				newBoard[i] = TokenO
				score := minimax(newBoard, 0, true)
				newBoard[i] = TokenEmpty

				if score < best {
					best = score
					bestMove = i
				}
			}
		}
	}

	return bestMove
}

func minimax(board *Board, depth int, isMax bool) int {
	score, isTie := checkWinner(board)

	if score == TokenX {
		return 10
	}

	if score == TokenO {
		return -10
	}

	if score == TokenEmpty && isTie {
		return 0
	}

	if isMax {
		// If this maximizer's or TokenX's move
		best := -1000
		for i := 0; i < 9; i++ {
			if board[i] == TokenEmpty {
				board[i] = TokenX
				best = maxScore(best, minimax(board, depth+1, false))
				board[i] = TokenEmpty
			}
		}
		return best - depth
	} else {
		// If this minimizer's or TokenO's move
		best := 1000
		for i := 0; i < 9; i++ {
			if board[i] == TokenEmpty {
				board[i] = TokenO
				best = minScore(best, minimax(board, depth+1, true))
				board[i] = TokenEmpty
			}
		}
		return best + depth
	}
}

func checkWinner(board *Board) (Token, bool) {
	winner := TokenEmpty

	// Horizontal
	if board[0] == board[1] && board[1] == board[2] && board[0] != TokenEmpty {
		if board[0] == TokenX {
			winner = TokenX
		} else {
			winner = TokenO
		}
	}
	if board[3] == board[4] && board[4] == board[5] && board[3] != TokenEmpty {
		if board[3] == TokenX {
			winner = TokenX
		} else {
			winner = TokenO
		}
	}
	if board[6] == board[7] && board[7] == board[8] && board[6] != TokenEmpty {
		if board[6] == TokenX {
			winner = TokenX
		} else {
			winner = TokenO
		}
	}

	// Vertical
	if board[0] == board[3] && board[3] == board[6] && board[0] != TokenEmpty {
		if board[0] == TokenX {
			winner = TokenX
		} else {
			winner = TokenO
		}
	}
	if board[1] == board[4] && board[4] == board[7] && board[1] != TokenO {
		if board[1] == TokenX {
			winner = TokenX
		} else {
			winner = TokenO
		}
	}
	if board[2] == board[5] && board[5] == board[8] && board[2] != TokenEmpty {
		if board[2] == TokenX {
			winner = TokenX
		} else {
			winner = TokenO
		}
	}

	// Diagonal
	if board[0] == board[4] && board[4] == board[8] && board[0] != TokenEmpty {
		if board[0] == TokenX {
			winner = TokenX
		} else {
			winner = TokenO
		}
	}
	if board[2] == board[4] && board[4] == board[6] && board[2] != TokenEmpty {
		if board[2] == TokenX {
			winner = TokenX
		} else {
			winner = TokenO
		}
	}

	openSpots := 0
	for i := 0; i < 9; i++ {
		if board[i] == TokenEmpty {
			openSpots++
		}
	}

	if winner == TokenEmpty && openSpots == 0 {
		return winner, true
	}

	return winner, false
}

func maxScore(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func minScore(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func isFirstMove(board *Board) bool {
	for _, el := range board {
		if el != TokenEmpty {
			return false
		}
	}
	return true
}
